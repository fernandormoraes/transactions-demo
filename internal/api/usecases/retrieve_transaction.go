package usecases

import (
	"strconv"
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/api/dto"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/persistence"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/remote"
	"github.com/fernandormoraes/transaction-demo/pkg/helpers"
)

type RetrieveTransaction interface {
	Run(date time.Time, currencyDesc string) (transactionsRetrieved []dto.TransactionRetrieveDto, err error)
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

func (u RetrieveTransactionUseCase) Run(date time.Time, currencyDesc string) (transactionsRetrieved []dto.TransactionRetrieveDto, err error) {
	transactions, err := u.transactionRepository.GetByDate(date)

	if err != nil {
		return make([]dto.TransactionRetrieveDto, 0), err
	}

	strDate := date.Format(helpers.LayoutDate)

	exchanges, err := u.remoteTreasury.FindAll(strDate, currencyDesc)

	if err != nil || len(exchanges) <= 0 {
		return make([]dto.TransactionRetrieveDto, 0), err
	}

	exchangeRate, err := strconv.ParseFloat(exchanges[0].ExchangeRate, 64)

	if err != nil {
		return nil, err
	}

	transactionsRetrieveDto := make([]dto.TransactionRetrieveDto, 0)

	for _, transaction := range transactions {
		transactionsRetrieveDto = append(transactionsRetrieveDto, dto.TransactionRetrieveDto{
			Description:     transaction.Description,
			Date:            transaction.Date,
			Amount:          transaction.Amount,
			ConvertedAmount: helpers.RoundFloat(transaction.Amount*exchangeRate, 2),
			ExchangeRate:    exchangeRate,
		})
	}

	return transactionsRetrieveDto, nil
}
