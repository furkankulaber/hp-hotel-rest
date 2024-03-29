package repository

import (
	"gorm.io/gorm"
	"hp-hotel-rest/internal/model"
)

type ReviewRepository interface {
	CreateReview(review *model.Review) (*model.Review, error)
}

type reviewRepository struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{DB: db}
}

func (r *reviewRepository) CreateReview(review *model.Review) (*model.Review, error) {
	if err := r.DB.Create(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (repo *reviewRepository) GetReviewsByHotelID(hotelID uint) ([]model.Review, error) {
	var reviews []model.Review
	if err := repo.DB.Where("hotel_id = ?", hotelID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}
