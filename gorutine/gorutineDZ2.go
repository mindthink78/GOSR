package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(transferPoint chan int) {
	randINT := rand.Intn(50)

	transferPoint <- randINT
}

func main() {
	temp := 0
	TransferPoint := make(chan int)

	go randomInt(TransferPoint)
	go randomInt(TransferPoint)
	go randomInt(TransferPoint)

	temp = <-TransferPoint
	time.Sleep(1 * time.Second)
	fmt.Println(temp)

	temp = <-TransferPoint
	time.Sleep(1 * time.Second)
	fmt.Println(temp)

	temp = <-TransferPoint
	time.Sleep(1 * time.Second)
	fmt.Println(temp)

}
