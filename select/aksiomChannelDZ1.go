package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandLetters() string {
	result := ""
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 5; i++ {
		letter := letters[rand.Intn(len(letters))] // Берет случайное число из кол-ва в letters и задает индекс,
		result += string(letter)                   // тем самым выбирается рандомная буква из массива букв letters
	}
	return result
}

func SleepTime() {
	RandTime := (3 + rand.Intn(5)) * 100
	time.Sleep(time.Duration(RandTime) * time.Millisecond)
}

func main() {
	transferPoint := make(chan string)

	PeopleValue := 1 + rand.Intn(10)
	fmt.Println("Интервьюер нашел:", PeopleValue, "людей")

	go func() {
		for i := 0; i < PeopleValue; i++ {

			SleepTime()
			transferPoint <- RandLetters()

		}
		close(transferPoint)
	}()

	for opinions := range transferPoint {
		fmt.Println(opinions)
	}

}
