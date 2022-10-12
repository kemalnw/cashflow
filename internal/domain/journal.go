package domain

import (
	"gorm.io/gorm"
	"time"
)

type Journal struct {
	gorm.Model
	Amount      uint      `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Source      Source    `json:"source,omitempty"`
	Category    Category  `json:"category,omitempty"`
}
