package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Cant Withdraw, Poor One")

// NewAccount create a new bank Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// String representation of the account
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account.\nHas: ", a.balance)
}

// Deposit on your account
func (a *Account) Deposit(amount int) {
	if amount > 0 {
		a.balance += amount
	}
}

// Withdraw x from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} else if amount > 0 {
		a.balance -= amount
	}
	return nil
}

// GetBalance of your  account
func (a Account) GetBalance() int {
	return a.balance
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// GetOwner of the account
func (a Account) GetOwner() string {
	return a.owner
}
