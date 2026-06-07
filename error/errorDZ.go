package main

import (
	"errors"
	"fmt"
)

func Calculator(a int, operator string, b int) (int, error) {
	if a > 1000 || a < -1000 || b > 1000 || b < -1000 {
		return 0, errors.New("Числа не подходят условию")
	}

	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":

		if b == 0 {
			return 0, errors.New("На ноль делить нельзя")
		}
		return a / b, nil

	default:
		return 0, errors.New("Неверный ввод")
	}
	return 0, nil
}

func main() {

	result, err := Calculator(5, "*", 5)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(result)

}
