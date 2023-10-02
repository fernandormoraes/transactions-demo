package api

import (
	"github.com/fernandormoraes/transaction-demo/internal/api/handlers"
	"github.com/fernandormoraes/transaction-demo/internal/api/usecases"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/persistence"
	"github.com/fernandormoraes/transaction-demo/internal/pkg/remote"
	"github.com/labstack/echo/v4"
)

func (s *Server) Router(e *echo.Echo) error {
	transactionRepository := persistence.NewTransactionRepository(s.db)
	remoteTreasury := remote.NewTreasuryRemote(s.httpClient)

	createTransaction := usecases.NewCreateTransaction(transactionRepository)
	retrieveTransaction := usecases.NewRetrieveTransactionUseCase(transactionRepository, remoteTreasury)

	transactionHandler := handlers.NewTransactionHandlers(createTransaction, retrieveTransaction)

	v1 := e.Group("/api/v1")

	transaction := v1.Group("/transactions")

	transaction.POST("", transactionHandler.PostTransaction)
	transaction.GET("", transactionHandler.RetrieveTransactions)

	return nil
}
