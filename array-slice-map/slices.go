package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Weight float64
}

func main() {
	intSlice := make([]int, 0)
	intSizeFocus := make([]int, 0, 5)

	fmt.Println(intSizeFocus)
	fmt.Println(len(intSizeFocus), cap(intSizeFocus))

	intSizeFocus = append(intSizeFocus, 12, 13, 14, 11, 34, 44)

	intSlice = append(intSlice, 10)

	fltSlice := []User{
		User{
			Name:   "Danil",
			Age:    21,
			Weight: 77,
		},
	}
	fmt.Println(fltSlice)
	fmt.Println("")

	fltSlice = append(fltSlice,
		User{
			Name:   "Nikita",
			Age:    19,
			Weight: 67,
		})

	fmt.Println(intSlice)
	fmt.Println(fltSlice)
	fmt.Println(intSizeFocus)
	fmt.Println(len(intSizeFocus), cap(intSizeFocus))

}
