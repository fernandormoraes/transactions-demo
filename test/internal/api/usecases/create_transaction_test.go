package usecases

import (
	"testing"

	"github.com/fernandormoraes/transaction-demo/internal/api/usecases"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
	"github.com/fernandormoraes/transaction-demo/test/mocks"
)

var createTransaction usecases.CreateTransaction = usecases.NewCreateTransaction(mocks.TransactionsRepository{})

func TestCreateTransaction(t *testing.T) {
	t.Run(`Given a new transaction
		   When creating a new transaction without errors
		   Then should'n return errors and return transaction
			`, func(t *testing.T) {
		transaction, err := createTransaction.Run(models.Transaction{})

		if err != nil {
			t.Errorf("Expected error to be nil, found: " + err.Error())
		}

		if transaction.Description != "Mocked transaction" {
			t.Errorf("Expected transaction to be mocked transaction")
		}
	})

	t.Run(`Given a new transaction
		   When creating a new transaction with an error
		   Then should return an error
			`, func(t *testing.T) {
		_, err := createTransaction.Run(models.Transaction{Description: "1337"})

		if err == nil {
			t.Errorf("Expected error, found nil")
		}
	})
}
