package main

import (
	"fmt"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Println("k:", k, "--- v:", v)
	}
	fmt.Println("")

}
func main() {
	http.HandleFunc("/header", headerHandler)

	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Сервер не запустился", err)
	}
}
