package main

import (
	"fmt"
	"time"
)

func main() {

	ChStr := make(chan string)
	ChInt := make(chan int)
	ChFl := make(chan float64)

	go func() {
		i := 1

		for {

			ChInt <- i
			i++

			time.Sleep(300 * time.Millisecond)

		}
	}()

	go func() {
		wordik := "Hello"

		for {
			ChStr <- wordik

			time.Sleep(1 * time.Second)
		}

	}()

	go func() {
		i := 1.11

		for {
			ChFl <- i
			i++

			time.Sleep(5 * time.Second)
		}

	}()

	for {

		select {
		case StrP := <-ChStr:
			fmt.Println("")
			fmt.Println("Строка:", StrP)
		case IntP := <-ChInt:
			fmt.Println("")
			fmt.Println("Инт:", IntP)
		case FlP := <-ChFl:
			fmt.Println("")
			fmt.Println("Флоат:", FlP)

		}

	}
}
