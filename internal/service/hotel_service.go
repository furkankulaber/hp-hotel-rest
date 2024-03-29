package service

import (
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/repository"
)

type HotelService interface {
	GetAllHotels() ([]model.HotelListResponse, error)
	GetHotelByID(id uint) (model.HotelDetailResponse, error)
	CreateHotel(hotel model.Hotel) (model.Hotel, error)
	UpdateHotel(hotel model.Hotel) (model.Hotel, error)
	DeleteHotel(id uint) error
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{repo: repo}
}

func (s *hotelService) GetAllHotels() ([]model.HotelListResponse, error) {
	hotels, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var hotelListResponse []model.HotelListResponse
	for _, hotel := range hotels {
		rating := hotel.CalculateRating()

		var photoURLs []string
		for _, photo := range hotel.Photos {
			photoURLs = append(photoURLs, photo.URL)
		}

		hotelListResponse = append(hotelListResponse, model.HotelListResponse{
			ID:    hotel.ID,
			Name:  hotel.Name,
			Stars: hotel.Stars,
			Type:  hotel.Type,
			Location: model.LocationResponse{
				City:     hotel.Location.City,
				District: hotel.Location.District,
				Address:  hotel.Location.Address,
			},
			Rating: rating,
			Photos: photoURLs,
		})
	}

	return hotelListResponse, nil
}

func (s *hotelService) GetHotelByID(id uint) (model.HotelDetailResponse, error) {
	hotel, err := s.repo.GetByID(id)
	if err != nil {
		return model.HotelDetailResponse{}, err
	}
	rating := hotel.CalculateRating()

	var roomResponses []model.RoomResponse
	for _, room := range hotel.Rooms {
		roomResponses = append(roomResponses, model.RoomResponse{
			Name:  room.Name,
			Price: room.Price,
		})
	}

	var reviewResponses []model.ReviewResponse
	for _, review := range hotel.Reviews {
		reviewResponses = append(reviewResponses, model.ReviewResponse{
			UserName: review.UserName,
			Rating:   review.Rating,
			Text:     review.Text,
		})
	}

	var photoURLs []string
	for _, photo := range hotel.Photos {
		photoURLs = append(photoURLs, photo.URL)
	}

	var amenities []string
	for _, amenity := range hotel.Amenities {
		amenities = append(amenities, amenity.Name)
	}

	return model.HotelDetailResponse{
		ID:          hotel.ID,
		Description: hotel.Description,
		Name:        hotel.Name,
		Stars:       hotel.Stars,
		Type:        hotel.Type,
		Rating:      rating,
		Location: model.LocationResponse{
			City:     hotel.Location.City,
			District: hotel.Location.District,
			Address:  hotel.Location.Address,
		},
		Amenities: amenities,
		Rooms:     roomResponses,
		Reviews:   reviewResponses,
		Photos:    photoURLs,
	}, nil
}

func (s *hotelService) CreateHotel(hotel model.Hotel) (model.Hotel, error) {
	return s.repo.Create(hotel)
}

func (s *hotelService) UpdateHotel(hotel model.Hotel) (model.Hotel, error) {
	return s.repo.Update(hotel)
}

func (s *hotelService) DeleteHotel(id uint) error {
	return s.repo.Delete(id)
}
