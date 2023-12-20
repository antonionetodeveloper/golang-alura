package checking_account

import (
	"bank/src/entities"
)

type CheckingAccount struct {
	Owner        string
	AgenceNumber int
	Account      int
	Ballance     entities.Money
}

func (person *CheckingAccount) Deposit(depositValue entities.Money) bool {
	if depositValue.Integer <= 0 || depositValue.Cents < 0 {
		return false
	}

	person.Ballance.Integer += depositValue.Integer
	person.Ballance.Cents += depositValue.Cents
	person.Ballance.CheckMoney()
	return true
}

func (person *CheckingAccount) Withdraw(withdrawValue entities.Money) bool {
	if withdrawValue.Integer > person.Ballance.Integer || withdrawValue.Integer <= 0 || withdrawValue.Cents < 0 {
		return false
	}

	if withdrawValue.Integer == person.Ballance.Integer {
		if withdrawValue.Cents > person.Ballance.Cents {
			return false
		}
	}

	person.Ballance.Integer -= withdrawValue.Integer
	person.Ballance.Cents -= withdrawValue.Cents
	person.Ballance.CheckMoney()
	return true
}

func (sender *CheckingAccount) Transfer(recipient *CheckingAccount, transferValue entities.Money) bool {
	if transferValue.Integer <= 0 || transferValue.Cents < 0 {
		return false
	}

	if sender.Ballance.Integer < transferValue.Integer {
		return false
	} else if sender.Ballance.Integer == transferValue.Integer {
		if sender.Ballance.Cents < transferValue.Cents {
			return false
		}
	}

	sender.Ballance.Integer -= transferValue.Integer
	sender.Ballance.Cents -= transferValue.Cents
	sender.Ballance.CheckMoney()
	recipient.Ballance.Integer += transferValue.Integer
	recipient.Ballance.Cents += transferValue.Cents
	recipient.Ballance.CheckMoney()
	return true
}
