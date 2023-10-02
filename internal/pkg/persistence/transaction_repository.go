package persistence

import (
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction models.Transaction) (*models.Transaction, error)
	GetByDate(date time.Time) (transactions []models.Transaction, err error)
}

type SqliteTransactionRepository struct {
	db *gorm.DB
}

var _ TransactionRepository = SqliteTransactionRepository{}

func NewTransactionRepository(db *gorm.DB) *SqliteTransactionRepository {
	return &SqliteTransactionRepository{db: db}
}

func (r SqliteTransactionRepository) Create(transaction models.Transaction) (*models.Transaction, error) {
	transaction.ID = uuid.New()

	tx := r.db.Create(&transaction)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &transaction, nil
}

func (r SqliteTransactionRepository) GetByDate(date time.Time) (transactions []models.Transaction, err error) {
	tx := r.db.Where("date >= ? AND date <= ?", date.Format("2006-01-02 00:00:01"), date.Format("2006-01-02 23:59:59")).Find(&transactions)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return transactions, nil
}
