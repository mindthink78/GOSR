package main

import (
	"fmt"
	"sync"
)

var slice []int

//var mtx sync.Mutex

func VotesAdd(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}
func main() {
	wg := &sync.WaitGroup{}

	wg.Add(5)
	go VotesAdd(wg)
	go VotesAdd(wg)
	go VotesAdd(wg)
	go VotesAdd(wg)
	go VotesAdd(wg)

	wg.Wait()

	mtx.Lock()
	fmt.Println(len(slice))
	mtx.Unlock()
}
