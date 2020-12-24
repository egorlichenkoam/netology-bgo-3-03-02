package card

import "01/pkg/money"

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	if transaction != nil {
		card.Transactions = append(card.Transactions, *transaction)
	}
}

func SumByMcc(transactions []Transaction, mccs []Mcc) money.Money {
	var result money.Money = 0
	transactions = filterTransactionsByMcc(transactions, mccs)
	for _, transaction := range transactions {
		result = result + transaction.Amount
	}
	return result
}

func filterTransactionsByMcc(transactions []Transaction, mccs []Mcc) []Transaction {
	result := make([]Transaction, 0)
	for _, transaction := range transactions {
		for _, mcc := range mccs {
			if transaction.Mcc == mcc {
				result = append(result, transaction)
			}
		}
	}
	return result
}

func TranslateMcc(code Mcc) string {
	result := "Категория не указана"
	value, ok := Mccs()[code]
	if ok {
		result = value
	}
	return result
}
