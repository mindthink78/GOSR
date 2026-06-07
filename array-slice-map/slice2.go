package main

import "fmt"

type Dog struct {
	Name    string
	Rating  int
	IfHappy bool
}

func main() {
	dogArr := []Dog{
		Dog{
			Name:    "Moza",
			Rating:  4,
			IfHappy: true,
		},
		Dog{
			Name:    "Perseus",
			Rating:  3,
			IfHappy: false,
		},
		Dog{
			Name:    "Dunai",
			Rating:  2,
			IfHappy: false,
		},
	}

	for i := 0; i < len(dogArr); i++ {
		fmt.Println(dogArr[i])
	}
	fmt.Println("")

	for i := 0; i < len(dogArr); i++ {
		if dogArr[i].IfHappy == true {
			dogArr[i].Rating += 1
		}
	}

	for i := 0; i < len(dogArr); i++ {
		fmt.Println(dogArr[i])
	}
	fmt.Println("")

}
