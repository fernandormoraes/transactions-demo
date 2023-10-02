package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey"`
	Description string
	Date        time.Time
	Amount      float64
}
