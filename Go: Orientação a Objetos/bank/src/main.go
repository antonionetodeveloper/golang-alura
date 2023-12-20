package main

import (
	account "bank/src/accounts"
	entities "bank/src/entities"
)

func main() {
	account1 := account.CheckingAccount{
		Owner:        "Antonio",
		AgenceNumber: 0567,
		Account:      123456,
		Ballance: entities.Money{
			Code:    "R$",
			Unit:    "BRL",
			Integer: 1380,
			Cents:   80,
		},
	}

	account2 := account.CheckingAccount{
		Owner:        "Antonio",
		AgenceNumber: 0567,
		Account:      123456,
		Ballance: entities.Money{
			Code:    "R$",
			Unit:    "BRL",
			Integer: 1380,
			Cents:   80,
		},
	}

	account1.Withdraw(entities.Money{
		Code:    "R$",
		Unit:    "BRL",
		Integer: 380,
		Cents:   90,
	})

	account2.Deposit(entities.Money{
		Code:    "R$",
		Unit:    "BRL",
		Integer: 380,
		Cents:   90,
	})

	account1.Transfer(&account2, entities.Money{
		Code:    "R$",
		Unit:    "BRL",
		Integer: 380,
		Cents:   90,
	})
}
