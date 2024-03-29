package repository

import (
	"gorm.io/gorm"
	"hp-hotel-rest/internal/model"
)

type ReviewRepository interface {
	CreateReview(review *model.Review) (*model.Review, error)
	GetReviewsByHotelID(hotelId uint) ([]model.Review, error)
	UpdateReview(review *model.Review) (*model.Review, error)
	DeleteReview(id uint) error
	FindReviewByID(id uint) (*model.Review, error)
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

func (r *reviewRepository) GetReviewsByHotelID(hotelID uint) ([]model.Review, error) {
	var reviews []model.Review
	if err := r.DB.Where("hotel_id = ?", hotelID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) UpdateReview(review *model.Review) (*model.Review, error) {
	if err := r.DB.Save(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *reviewRepository) DeleteReview(id uint) error {
	if err := r.DB.Delete(&model.Review{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) FindReviewByID(id uint) (*model.Review, error) {
	var review model.Review
	result := r.DB.First(&review, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &review, nil
}
