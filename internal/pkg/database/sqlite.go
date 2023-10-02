package database

import (
	"github.com/fernandormoraes/transaction-demo/internal/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("transactions.db"), &gorm.Config{})

	db.AutoMigrate(&models.Transaction{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
