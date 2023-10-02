package mocks

import (
	"errors"
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
)

type TransactionsRepository struct{}

func (mock TransactionsRepository) Create(transaction models.Transaction) (*models.Transaction, error) {
	if transaction.Description == "1337" {
		return nil, errors.New("Error")
	}

	return &models.Transaction{
		Description: "Mocked transaction",
		Amount:      2.00,
	}, nil
}

func (mock TransactionsRepository) GetByDate(date time.Time) (transactions []models.Transaction, err error) {
	listTransactions := make([]models.Transaction, 0)

	listTransactions = append(listTransactions, models.Transaction{
		Description: "Mocked transaction",
		Amount:      2.00,
	})

	return listTransactions, nil
}
