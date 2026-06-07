package main

import (
	"fmt"
	"time"
)

func Counter(x int) {
	for i := 1; i < 6; i++ {
		fmt.Printf("Я горутина %v делаю вывод на экран %v раз\n", x, i)
		time.Sleep(time.Second)
	}

}

func main() {
	go Counter(1)
	go Counter(2)
	go Counter(3)

	time.Sleep(10 * time.Second)

}
