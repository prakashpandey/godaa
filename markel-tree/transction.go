package main

import (
	"time"

	"github.com/prakashpandey/godaa/markel-tree/transaction"
)

func generateTransactions() {
	for i := 0; i < 11; i++ {
		transaction := func() *transaction.Transaction {
			transac := transaction.New()
			transac.From = "account-1"
			transac.To = "account-2"
			transac.Timestamp = time.Now().UnixNano()
			transac.Out = []transaction.Detail{
				{
					Address: "account-1",
					Amount:  0.2,
					Type:    transaction.OUT,
				},
			}
			return transac
		}
		addTransaction(transaction())
	}
}

func addTransaction(t *transaction.Transaction) {
	transactions = append(transactions, t)
}
