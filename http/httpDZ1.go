package main

import (
	"fmt"
	"net/http"
)

func DogHandler(w http.ResponseWriter, r *http.Request) {
	str := "Я собака и я говорю 'Гав'"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP запроса произошла ошибка")
	} else {
		fmt.Println("HTTP запрос корректно обработан")
	}

}
func CatHandler(w http.ResponseWriter, r *http.Request) {
	str := "Я кошка и я говрою 'Мяу'"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP запроса произошла ошибка")
	} else {
		fmt.Println("HTTP запрос корректно обработан")
	}

}
func CowHandler(w http.ResponseWriter, r *http.Request) {
	str := "Я корова и я говорю 'Myyy'"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP запроса произошла ошибка")
	} else {
		fmt.Println("HTTP запрос корректно обработан")
	}

}

func main() {
	http.HandleFunc("/dog", DogHandler)
	http.HandleFunc("/cat", CatHandler)
	http.HandleFunc("/cow", CowHandler)
	fmt.Println("Запускаю HTTP сервер")
	err := http.ListenAndServe(":9091", nil)

	if err != nil {
		fmt.Println("Произошла ошибку")
	}
	fmt.Println("Сервер выключен")
}
