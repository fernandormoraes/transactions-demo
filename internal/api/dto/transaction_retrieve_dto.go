package dto

import "time"

type TransactionRetrieveDto struct {
	Description     string
	Date            time.Time
	Amount          float64
	ConvertedAmount float64
	ExchangeRate    float64
}
