package model

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	URL     string
	HotelID uint
	Hotel   Hotel
}
