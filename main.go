package main

import (
	"fmt"

	"github.com/LeSpank/bankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("mohamed")
	fmt.Println(account)
	account.Deposit(200)
	fmt.Println(account.GetBalance())
	err := account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(accounts.NewAccount("said"))
	fmt.Println(account.GetOwner())

}
