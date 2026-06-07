package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func farmer(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	randTime := 500 + rand.Intn(501)

	fmt.Println("я огородник", i, "Поливаю грядки")
	time.Sleep(time.Duration(randTime) * time.Millisecond)

	fmt.Println("я огородник", i, ", закончил рабочий день!")
}

func main() {
	wg := &sync.WaitGroup{}

	randValueGuys := 1 + rand.Intn(5)

	for i := 1; i < randValueGuys; i++ {
		wg.Add(1)
		go farmer(wg, i)
	}

	wg.Wait()
	fmt.Println("Рабочий день окончен!")
}
