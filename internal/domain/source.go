package domain

import "gorm.io/gorm"

type Source struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
