package model

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	City     string
	District string
	Address  string
	Hotels   []Hotel `gorm:"foreignKey:LocationID"`
}

// DTO
type LocationResponse struct {
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}
