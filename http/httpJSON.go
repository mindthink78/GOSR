package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
}

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

type HttpResponse struct {
	Money          int       `json:"mmoney"`
	PaymentHistory []Payment `json:"pHistory"`
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Вот это все можно заменить тем, что сверху

	//httpRequestBody, err := io.ReadAll(r.Body)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//
	//if err := json.Unmarshal(httpRequestBody, &payment); err != nil {
	//	fmt.Println("err", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}

	paymentHistory = append(paymentHistory, payment)

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("money:", money)
	fmt.Println("payment history:", paymentHistory)

	mtx.Unlock()

}

func main() {
	http.HandleFunc("/pay", payHandler)

	fmt.Println("HTTP server start")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы сервера", err)
	}
}
