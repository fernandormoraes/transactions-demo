package handlers

import (
	"net/http"
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/api/usecases"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
	"github.com/fernandormoraes/transaction-demo/pkg/helpers"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TransactionHandlers struct {
	createTransaction   usecases.CreateTransaction
	retrieveTransaction usecases.RetrieveTransaction
}

func NewTransactionHandlers(
	createTransaction usecases.CreateTransaction,
	retrieveTransaction usecases.RetrieveTransaction) *TransactionHandlers {
	return &TransactionHandlers{
		createTransaction:   createTransaction,
		retrieveTransaction: retrieveTransaction,
	}
}

type (
	CreateTransactionDTO struct {
		Description string  `json:"description" validate:"required,gt=0,lt=51"`
		Date        string  `json:"date" validate:"required"`
		Amount      float64 `json:"amount"`
	}

	CreateTransactionValidator struct {
		Validator *validator.Validate
	}
)

// Validate implements echo.Validator.
func (cv *CreateTransactionValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (h TransactionHandlers) PostTransaction(c echo.Context) error {
	var createTransaction CreateTransactionDTO

	err := c.Bind(&createTransaction)

	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	validator := &CreateTransactionValidator{Validator: validator.New()}

	validator.Validator.Struct(&createTransaction)

	err = c.Validate(&createTransaction)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	dateTime, err := time.Parse("2006-01-02 15:04:05", createTransaction.Date)

	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	newTransaction := models.Transaction{
		ID:          uuid.New(),
		Description: createTransaction.Description,
		Amount:      helpers.RoundFloat(createTransaction.Amount, 2),
		Date:        dateTime,
	}

	createdTransaction, err := h.createTransaction.Run(newTransaction)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, createdTransaction)
}

func (h TransactionHandlers) RetrieveTransactions(c echo.Context) error {
	date := c.QueryParam("transactionDate")
	currencyDesc := c.QueryParam("currencyDesc")

	dateTime, err := time.Parse(helpers.LayoutDate, date)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	transactions, err := h.retrieveTransaction.Run(dateTime, currencyDesc)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transactions)
}
