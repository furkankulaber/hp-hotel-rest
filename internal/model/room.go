package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name    string
	Price   float64
	HotelID uint
	Hotel   Hotel
}

// DTO
type RoomResponse struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
