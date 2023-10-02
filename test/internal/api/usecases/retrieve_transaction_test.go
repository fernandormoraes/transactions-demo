package usecases

import (
	"testing"
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/api/usecases"
	"github.com/fernandormoraes/transaction-demo/test/mocks"
)

var retrieveTransaction usecases.RetrieveTransaction = usecases.NewRetrieveTransactionUseCase(mocks.TransactionsRepository{}, mocks.TreasuryRemote{})

func TestRetrieveTransaction(t *testing.T) {
	t.Run(`Given transaction date and country currency
		   When found exchanges and transactions
		   Then should return transaction with converted amount
			`, func(t *testing.T) {
		transactions, err := retrieveTransaction.Run(time.Now(), "Brazil-Real")

		if err != nil {
			t.Errorf("Expected error to be nil, found: " + err.Error())
		}

		if transactions[0].ConvertedAmount != 10.00 {
			t.Errorf("Expected transaction to be mocked transaction with converted amount")
		}
	})

	t.Run(`Given a new transaction
		   When creating a new transaction with an error
		   Then should return an error
			`, func(t *testing.T) {
		_, err := retrieveTransaction.Run(time.Now(), "1337")

		if err == nil {
			t.Errorf("Expected error, found nil")
		}
	})
}
