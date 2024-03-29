package repository

import (
	"gorm.io/gorm"
	"hp-hotel-rest/internal/model"
)

type HotelRepository interface {
	GetAll() ([]model.Hotel, error)
	GetByID(id uint) (model.Hotel, error)
	Create(hotel model.Hotel) (model.Hotel, error)
	Update(hotel model.Hotel) (model.Hotel, error)
	Delete(id uint) error
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
	result := r.DB.Preload("Location").Preload("Photos").Preload("Reviews").First(&hotel, id)
	return hotel, result.Error
}

func (r *hotelRepository) Create(hotel model.Hotel) (model.Hotel, error) {
	result := r.DB.Create(&hotel)
	return hotel, result.Error
}

func (r *hotelRepository) Update(hotel model.Hotel) (model.Hotel, error) {
	result := r.DB.Save(&hotel)
	return hotel, result.Error
}

func (r *hotelRepository) Delete(id uint) error {
	result := r.DB.Delete(&model.Hotel{}, id)
	return result.Error
}