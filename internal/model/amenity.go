package model

import (
	"gorm.io/gorm"
)

type Amenity struct {
	gorm.Model
	Name string
}

// DTO
type AmenityResponse struct {
	Name string `json:"name"`
}
