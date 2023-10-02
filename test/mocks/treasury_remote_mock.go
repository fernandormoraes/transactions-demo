package mocks

import (
	"errors"

	"github.com/fernandormoraes/transaction-demo/internal/pkg/remote"
)

type TreasuryRemote struct{}

func (mock TreasuryRemote) FindAll(transactionDate string, currencyDesc string) ([]remote.Exchange, error) {
	if currencyDesc == "1337" {
		return nil, errors.New("Error")
	}

	listExchange := make([]remote.Exchange, 0)

	listExchange = append(listExchange, remote.Exchange{
		ExchangeRate: "5.00",
		Country:      "Brazil",
		Currency:     "Brazil-Real",
		RecordDate:   "2023-10-02",
	})

	return listExchange, nil
}
