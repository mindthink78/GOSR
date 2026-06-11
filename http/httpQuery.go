package main

import (
	"fmt"
	"net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	firstParam := r.URL.Query().Get("first")
	secondParam := r.URL.Query().Get("second")

	fmt.Println("first:", firstParam)
	fmt.Println("second:", secondParam)
}

func main() {
	http.HandleFunc("/query", queryHandler)

	fmt.Println("SERVER STARTED")

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Server crashed")
		return
	}
}
