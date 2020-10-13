package card

type Transaction struct {
	Id     int64
	Amount int64
	Date   string
	Mcc    string
	Status string
}
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
	card.Transactions = append(card.Transactions, *transaction)
}

func SumByMCC(transactions []Transaction, mcc []string) int64 {
	total := int64(0)
	for _, transaction := range transactions {
		for _, code := range mcc {
			if transaction.Mcc == code {
				total += transaction.Amount
			}
		}
	}
	return total
}
func TranslateMCC(code string) string {
	mccText := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"0001": "Пополнения",
	}
	value, ok := mccText[code]
	if ok == true {
		return value
	}
	return "Категория не указана"
}
