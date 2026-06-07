package main

import "fmt"

func main() {
	parking := map[string]int{
		"A4": 200,
		"B2": 499,
		"C4": 399,
		"M8": 800,
		"N1": 780,
		"K8": 990,
		"Y0": 1100,
	}

	for k, _ := range parking {
		if parking[k] < 500 {
			fmt.Println(parking[k])

		}
		if parking[k] > 900 {
			discount := float64(parking[k]) * 0.10
			parking[k] -= int(discount)
			fmt.Println(parking[k])
		}
	}
}
