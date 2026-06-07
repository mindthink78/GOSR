package main

import (
	"fmt"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		fmt.Println("Method allowed")
		msg := "Method allowed"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("err:", err)
		}
	}

}

func main() {

	http.HandleFunc("/method", methodHandler)

	fmt.Println("HTTP SERVER ONLINE")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("err:", err)

	}

}
