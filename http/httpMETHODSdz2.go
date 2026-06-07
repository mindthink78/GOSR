package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type MessageStorage struct {
	msg map[int]string
	mtx sync.RWMutex
}

func NewMessageStorage() MessageStorage {
	return MessageStorage{
		msg: make(map[int]string),
		mtx: sync.RWMutex{},
	}
}

func (m *MessageStorage) messageHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("FAIL")
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	m.mtx.Lock()
	ID := 1000 + rand.Intn(1000)
	m.msg[ID] = httpRequestBodyString
	fmt.Println("Saved message:", m.msg[ID])
	m.mtx.Unlock()

	_, err = w.Write([]byte("Message saved"))
	if err != nil {
		fmt.Println("FAIL")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//m.mtx.RLock()
	//for id, msg := range m.msg {
	//	fmt.Println(id, msg)
	//}
	//m.mtx.RUnlock()
}

func (m *MessageStorage) deleteHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("FAIl")
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	id, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("FAIL")
		http.Error(w, "404 not found", http.StatusBadRequest)
	}

	_, ok := m.msg[id]
	if !ok {
		msg := "По такому айди ничего нет"
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Не получилось вывести у клиента")
			http.Error(w, "Чето не получилось", http.StatusBadRequest)
		}
	} else {
		m.mtx.Lock()
		delete(m.msg, id)
		m.mtx.Unlock()
	}

}

func (m MessageStorage) savedMessageHandler(w http.ResponseWriter, r *http.Request) {
	m.mtx.RLock()
	for id, msg := range m.msg {
		fmt.Println(id, msg)
	}
	m.mtx.RUnlock()
}

func main() {
	m := NewMessageStorage()
	http.HandleFunc("/message", m.messageHandler)
	http.HandleFunc("/delete", m.deleteHandler)
	http.HandleFunc("/savedmessage", m.savedMessageHandler)

	fmt.Println("Сервер запустился")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error", err)
	}
}
