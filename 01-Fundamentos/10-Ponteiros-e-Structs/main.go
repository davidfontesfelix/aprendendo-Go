package main

import "fmt"

type CurrentAccount struct {
	balance int
}

func NewAccount() *CurrentAccount {
	return &CurrentAccount{balance: 0}
}

func (currentAccount *CurrentAccount) simulate(value int) int {
	currentAccount.balance += value
	println(currentAccount.balance)
	return currentAccount.balance
}

func main() {
	account := CurrentAccount{
		balance: 100,
	}
	account.simulate(200)

	fmt.Println(account.balance)
}
