package repository

import (
	"gorm.io/gorm"
	"hp-hotel-rest/internal/model"
)

type ReviewRepository interface {
	Create(review model.Review) (model.Review, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(review model.Review) (model.Review, error) {
	// Burada veritabanına review oluşturma işlemini gerçekleştir
}
