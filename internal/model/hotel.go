package model

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string
	Description string
	Stars       int
	Type        string
	Rating      float64 `gorm:"-"`
	LocationID  uint
	Location    Location   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amenities   []*Amenity `gorm:"many2many:hotel_amenities;"`
	Rooms       []Room     `gorm:"foreignKey:HotelID"`
	Reviews     []Review   `gorm:"foreignKey:HotelID"`
	Photos      []Photo    `gorm:"foreignKey:HotelID"`
}

func (h *Hotel) CalculateRating() float64 {
	var totalRating int
	for _, review := range h.Reviews {
		totalRating += review.Rating
	}
	if len(h.Reviews) > 0 {
		return float64(totalRating) / float64(len(h.Reviews))
	}
	return 0
}

// DTO

type HotelDetailResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Stars       int              `json:"stars"`
	Type        string           `json:"type"`
	Location    LocationResponse `json:"location"`
	Rating      float64          `json:"rating"`
	Amenities   []string         `json:"amenities"`
	Rooms       []RoomResponse   `json:"rooms"`
	Reviews     []ReviewResponse `json:"reviews"`
	Photos      []string         `json:"photos"`
}

type HotelListResponse struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Stars    int              `json:"stars"`
	Type     string           `json:"type"`
	Location LocationResponse `json:"location"`
	Rating   float64          `json:"rating"`
	Photos   []string         `json:"photos"`
}
