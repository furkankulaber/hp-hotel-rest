package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserName  string
	UserEmail string
	Score     int
	Text      string
	HotelID   uint
}

// DTO
type ReviewResponse struct {
	UserName string `json:"user_name"`
	Score    int    `json:"score"`
	Text     string `json:"text"`
}
