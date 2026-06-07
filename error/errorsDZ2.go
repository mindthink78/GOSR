package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

type Account struct {
	Name    string
	Balance int
}

func CashOut(a *Account, value int) error {
	if a.Balance < value {
		return errors.New("Недостаточно средств")
	}
	a.Balance -= value
	fmt.Println("Вы вывели:", value)
	return nil
}

func BalanceInfo(a Account) {
	fmt.Println(a.Balance)
}

func Pay(a *Account, value int, item string) error {
	z := rand.IntN(3)
	if a.Balance < value {
		return errors.New("Недостаточно средств")
	} else if z == 2 {
		return errors.New("Оплата не прошла, повторите")
	}

	a.Balance -= value
	fmt.Println("Вы приобрели:", item, "за", value, "руб")
	return nil
}

func main() {
	User1 := Account{
		Name:    "Sanek",
		Balance: 4345,
	}

	for {
		cost := rand.IntN(600)

		err := Pay(&User1, cost, "iphine")
		if err != nil {
			fmt.Println("Оплата не была произведена! Причина:", err.Error())
		} else {
			fmt.Println("Оплата успешно произведена!")
		}
		BalanceInfo(User1)
		time.Sleep(1 * time.Second)

	}

}
