package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	FirChannelPoint := make(chan int)
	SecChannelPoint := make(chan int)

	go func() {
		value := (rand.Intn(10) + 1) * 100
		time.Sleep(time.Duration(value) * time.Millisecond)
		fmt.Println("Горутина спала", value, "ms")

		FirChannelPoint <- value

	}()

	go func() {
		value := (rand.Intn(10) + 1) * 100
		time.Sleep(time.Duration(value) * time.Millisecond)
		fmt.Println("Горутина спала", value, "ms")

		SecChannelPoint <- value

	}()

	time.Sleep(500 * time.Millisecond)

	select {
	case FirstCh := <-FirChannelPoint:
		fmt.Println("Отработал первый канал")
		fmt.Println("value:", FirstCh)
	case SecondCh := <-SecChannelPoint:
		fmt.Println("Отработал второй канал")
		fmt.Println("value:", SecondCh)
	default:
		fmt.Println("Отработал default, никакие значения каналов, не были переданы")
	}
}
