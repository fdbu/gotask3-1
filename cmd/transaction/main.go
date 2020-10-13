package main

import (
	"fmt"
	"gotask3-1/pkg/card"
)

func main() {
	t := &card.Transaction{
		Id:     3,
		Amount: 2000.00,
		Date:   "2020-07-19 21:00:00 +0000 UTC",
		Mcc:    "0001",
		Status: "done",
	}
	master := &card.Card{
		Id:       1,
		Issuer:   "MasterCard",
		Balance:  100_000_00,
		Currency: "RUB",
		Number:   "0001",
		Icon:     "https://mpng.subpng.com/20171216/dcc/mastercard-icon-png-5a3556c6e81b34.5328243515134450629507.jpg",
		Transactions: []card.Transaction{
			{
				Id:     1,
				Amount: 73555,
				Date:   "2020-07-16 20:55:46 +0000 UTC",
				Mcc:    "5411",
				Status: "in processing",
			},
			{
				Id:     2,
				Amount: 120391,
				Date:   "2020-07-18 13:54:40 +0000 UTC",
				Mcc:    "5812",
				Status: "done",
			},
		},
	}
	card.AddTransaction(master, t)
	mcc := []string{"5411", "5812"}
	fmt.Println(card.SumByMCC(master.Transactions, mcc))
}
