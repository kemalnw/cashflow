package domain

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name    string `json:"name"`
	Balance uint   `json:"balance"`
}
