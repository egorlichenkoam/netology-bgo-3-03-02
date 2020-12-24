package card

import "01/pkg/money"

type Transaction struct {
	Id       int64
	Amount   money.Money
	Datetime int64
	Mcc      Mcc
	Status   string
}
