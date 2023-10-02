package usecases

import (
	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/persistence"
)

type CreateTransaction interface {
	Run(transaction models.Transaction) (*models.Transaction, error)
}

type CreateTransactionUseCase struct {
	transactionRepository persistence.TransactionRepository
}

func NewCreateTransaction(transactionRepository persistence.TransactionRepository) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{transactionRepository: transactionRepository}
}

func (u CreateTransactionUseCase) Run(transaction models.Transaction) (*models.Transaction, error) {
	newTransaction, err := u.transactionRepository.Create(transaction)

	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}
