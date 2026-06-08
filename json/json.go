package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	FullName string  `json:"fullName"`
	Address  string  `json:"address"`
	Age      int     `json:"age"`
	Married  bool    `json:"married"`
	Height   float64 `json:"height"`
}

type UserStorage struct {
	Users []User
	mtx   sync.RWMutex
}

type Server struct {
	storage *UserStorage
}

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if r.Method != "POST" {
		fmt.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.storage.mtx.Lock()
	s.storage.Users = append(s.storage.Users, user)
	fmt.Println(user)
	defer s.storage.mtx.Unlock()

	msg := "Данные успешно сохранены"

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		fmt.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	s.storage.mtx.RLock()
	b, err := json.Marshal(s.storage.Users)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	s.storage.mtx.RUnlock()
	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("Данные успешно отправлены клиенту")

}

func main() {
	storage := &UserStorage{
		Users: make([]User, 0),
		mtx:   sync.RWMutex{},
	}
	server := &Server{
		storage: storage,
	}

	http.HandleFunc("/user/create/", server.createUserHandler)
	http.HandleFunc("/user/get/", server.getUserHandler)

	fmt.Println("SERVER STARTED")

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Server crashed")
	}
}
