package repository

import (
	"gorm.io/gorm"
	"hp-hotel-rest/internal/model"
)

type HotelRepository interface {
	GetAll() ([]model.Hotel, error)
	GetByID(id uint) (model.Hotel, error)
}

type hotelRepository struct {
	DB *gorm.DB
}

func NewHotelRepository(db *gorm.DB) HotelRepository {
	return &hotelRepository{DB: db}
}

func (r *hotelRepository) GetAll() ([]model.Hotel, error) {
	var hotels []model.Hotel
	result := r.DB.Preload("Location").Preload("Photos").Preload("Reviews").Find(&hotels)
	return hotels, result.Error
}

func (r *hotelRepository) GetByID(id uint) (model.Hotel, error) {
	var hotel model.Hotel
	result := r.DB.
		Preload("Location").
		Preload("Photos").
		Preload("Reviews").
		Preload("Amenities").
		Preload("Rooms").
		First(&hotel, id)
	return hotel, result.Error
}
