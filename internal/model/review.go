package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserName  string
	UserEmail string
	Rating    int
	Text      string
	HotelID   uint
}

// DTO
type ReviewResponse struct {
	UserName string `json:"user_name"`
	Rating   int    `json:"rating"`
	Text     string `json:"text"`
}

type CreateReviewRequest struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Rating    int    `json:"rating"`
	Text      string `json:"text"`
}

type UpdateReviewRequest struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Rating    int    `json:"rating"`
	Text      string `json:"text"`
}
