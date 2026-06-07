package main

import (
	"context"
	"fmt"
	"time"
)

func Child(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Функция Child завершилась")
			return
		default:
			fmt.Println("Child")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func Middle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Функция Middle завершилась")
			return
		default:
			fmt.Println("Middle")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func Parent(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Функция Parent завершилась")
			return
		default:
			fmt.Println("Parent")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	middleContext, middleCancel := context.WithCancel(parentContext)
	childContext, childCancel := context.WithCancel(middleContext)

	go Child(childContext)
	go Middle(middleContext)
	go Parent(parentContext)

	time.Sleep(1 * time.Second)
	middleCancel()
	time.Sleep(200 * time.Millisecond)
	parentCancel()

	time.Sleep(1 * time.Second)
	childCancel()

}
