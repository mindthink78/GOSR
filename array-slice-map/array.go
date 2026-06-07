package main

import "fmt"

func FullArr[T any](arr [3]T) { // T это дженерик, который позволяет, сделать функцию универсальной
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

//func FullArr(arr [3]int) {
//	for i := 0; i < len(arr); i++ {
//		fmt.Println(arr[i])
//	}
//}

func main() {
	arrInt := [3]int{11, 22, 33}
	arrStr := [3]string{"первый", "второй", "третий"}
	arrFlt := [3]float64{1.1, 2.2, 3.3}

	fmt.Println(arrInt[0])
	fmt.Println(arrStr[1])
	fmt.Println(arrFlt[2])
	fmt.Println("")

	FullArr(arrInt)
	FullArr(arrStr)
	FullArr(arrFlt)
	fmt.Println("")

	arrInt[0] = 21
	arrStr[1] = "probe"
	arrFlt[2] = 6.7
	FullArr(arrInt)
	FullArr(arrStr)
	FullArr(arrFlt)

}
