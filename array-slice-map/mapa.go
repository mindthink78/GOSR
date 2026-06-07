package main

import "fmt"

func main() {
	database := make(map[string]int)

	database["beeline"] = 909
	database["megafon"] = 922
	database["tele2"] = 921
	database["MTC"] = 913

	fmt.Println(database["beeline"])

	parking := map[int]bool{
		101: true,
		102: false,
		103: false,
		104: true,
		105: true,
	}
	fmt.Println(parking[102])
	fmt.Println("")

	parking[100] = true

	fmt.Println(parking[102])
	fmt.Println("")

	free, ok := parking[102]

	if !ok {
		fmt.Println("Такого места нет")
		fmt.Println("")
		return
	}

	if free {
		fmt.Println("Место свободно")
		fmt.Println("")
	} else {
		fmt.Println("Место занято")
		fmt.Println("")
	}

	fmt.Println(free)

}
