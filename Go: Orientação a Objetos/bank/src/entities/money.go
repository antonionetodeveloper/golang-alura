package entities

type Money struct {
	Code    string
	Unit    string
	Integer int
	Cents   int
}

func (money *Money) CheckMoney() {
	if money.Cents >= 100 {
		money.Integer++
		money.Cents -= 100
	}

	if money.Cents < 0 {
		money.Integer--
		money.Cents += 100
	}
}
