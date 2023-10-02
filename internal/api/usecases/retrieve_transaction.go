package usecases

import (
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/persistence"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/remote"
)

type RetrieveTransaction interface {
	Run(date time.Time) (transactions []models.Transaction, err error)
}

type RetrieveTransactionUseCase struct {
	transactionRepository persistence.TransactionRepository
	remoteTreasury        remote.TreasuryRemote
}

func NewRetrieveTransactionUseCase(
	transactionRepository persistence.TransactionRepository,
	remoteTreasury remote.TreasuryRemote) *RetrieveTransactionUseCase {
	return &RetrieveTransactionUseCase{
		transactionRepository: transactionRepository,
		remoteTreasury:        remoteTreasury,
	}
}

func (u RetrieveTransactionUseCase) Run(date time.Time) (transactions []models.Transaction, err error) {
	transactions, err = u.transactionRepository.GetByDate(date)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
