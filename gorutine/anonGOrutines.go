package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("GORUTINE1")
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("GORUTINE2")
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("GORUTINE3")
		}
	}()

	time.Sleep(3 * time.Second)
}
