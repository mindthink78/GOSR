package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	products := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите команду:")
		scanner.Scan()

		text := scanner.Text()
		arr := strings.Fields(text)

		if len(arr) == 0 {
			fmt.Println("Вы ничего не ввели")
			return
		}
		cmd := arr[0]
		amount := 0
		if len(arr) >= 3 {
			amount, _ = strconv.Atoi(arr[2])
		}
		eda := arr[1]

		if cmd == "добавить" {
			fmt.Print("Вы добавили: ", eda)
			products[eda] += amount
			fmt.Println("")
			fmt.Println(products[eda])

			continue
		} else if cmd == "удалить" {
			products[eda] -= amount
			fmt.Println("Вы удалили:", amount, eda)
			fmt.Println(products[eda])
			continue
		} else if cmd == "help" {
			fmt.Println("Команда получить")
			fmt.Println("Команда добавить")
			fmt.Println("Команда удалить")
			continue
		} else if cmd == "exit" {
			return
		} else if cmd == "получить" {
			fmt.Println(products[eda])
			continue
		} else {
			fmt.Println("Неизвестная команда")
			continue
		}

	}
}
