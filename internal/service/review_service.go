package service

import (
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/repository"
)

type ReviewService interface {
	AddReview(request *model.CreateReviewRequest, hotel *model.Hotel) (*model.ReviewResponse, error)
	GetReviewsByHotelID(uint) ([]model.ReviewResponse, error)
	UpdateReview(id uint, review *model.UpdateReviewRequest) (*model.ReviewResponse, error)
	DeleteReview(id uint) error
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) AddReview(request *model.CreateReviewRequest, hotel *model.Hotel) (*model.ReviewResponse, error) {
	review := &model.Review{
		UserName:  request.UserName,
		UserEmail: request.UserEmail,
		Rating:    request.Rating,
		Text:      request.Text,
		HotelID:   hotel.ID,
	}

	createdReview, err := s.repo.CreateReview(review)
	if err != nil {
		return nil, err
	}

	response := &model.ReviewResponse{
		UserName: createdReview.UserName,
		Rating:   createdReview.Rating,
		Text:     createdReview.Text,
	}

	return response, nil
}

func (s *reviewService) GetReviewsByHotelID(hotelID uint) ([]model.ReviewResponse, error) {
	reviews, err := s.repo.GetReviewsByHotelID(hotelID)
	if err != nil {
		return nil, err
	}

	var responses []model.ReviewResponse
	for _, review := range reviews {
		responses = append(responses, model.ReviewResponse{
			UserName: review.UserName,
			Rating:   review.Rating,
			Text:     review.Text,
		})
	}

	return responses, nil
}

func (s *reviewService) UpdateReview(id uint, req *model.UpdateReviewRequest) (*model.ReviewResponse, error) {
	review, err := s.repo.FindReviewByID(id)
	if err != nil {
		return nil, err
	}

	review.UserName = req.UserName
	review.UserEmail = req.UserEmail
	review.Rating = req.Rating
	review.Text = req.Text

	updatedReview, _ := s.repo.UpdateReview(review)

	return &model.ReviewResponse{
		UserName: updatedReview.UserName,
		Rating:   updatedReview.Rating,
		Text:     updatedReview.Text,
	}, nil

}

func (s *reviewService) DeleteReview(id uint) error {
	return s.repo.DeleteReview(id)
}
