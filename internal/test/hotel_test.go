package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hp-hotel-rest/internal/handler"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/pkg/utils"
)

type MockHotelService struct {
	mock.Mock
}

func (m *MockHotelService) GetAllHotels() ([]model.HotelListResponse, error) {
	args := m.Called()
	return args.Get(0).([]model.HotelListResponse), args.Error(1)
}

func (m *MockHotelService) GetHotelByID(id uint) (model.HotelDetailResponse, error) {
	args := m.Called(id)
	return args.Get(0).(model.HotelDetailResponse), args.Error(1)
}

func TestHotelHandler_GetAllHotels(t *testing.T) {
	mockService := new(MockHotelService)
	hotelHandler := handler.NewHotelHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		expectedHotels := []model.HotelListResponse{
			{ID: 1, Name: "Hotel A", Stars: 5, Type: "Luxury"},
			{ID: 2, Name: "Hotel B", Stars: 4, Type: "Standard"},
		}
		mockService.On("GetAllHotels").Return(expectedHotels, nil)

		app := fiber.New()
		app.Get("/hotels", hotelHandler.GetAllHotels)

		req := httptest.NewRequest(http.MethodGet, "/hotels", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response utils.APIResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Hotels fetched successfully", response.Message)

		dataBytes, err := json.Marshal(response.Data)
		assert.NoError(t, err)

		var hotels []model.HotelListResponse
		err = json.Unmarshal(dataBytes, &hotels)
		assert.NoError(t, err)

		mockService.AssertExpectations(t)
	})

}

func TestHotelHandler_GetHotelByID(t *testing.T) {
	mockService := new(MockHotelService)
	hotelHandler := handler.NewHotelHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		expectedHotel := model.HotelDetailResponse{
			ID:          1,
			Name:        "Hotel A",
			Description: "Luxury hotel with great amenities",
			Stars:       5,
			Type:        "Luxury",
			Location:    model.LocationResponse{City: "City A"},
			Rating:      4.5,
			Amenities:   []string{"Pool", "Gym", "Spa"},
			Rooms:       []model.RoomResponse{{Name: "Suite", Price: 200}, {Name: "Standard", Price: 100}},
			Reviews:     []model.ReviewResponse{{UserName: "User1", Rating: 5, Text: "Great experience"}},
			Photos:      []string{"photo1.jpg", "photo2.jpg"},
		}
		mockService.On("GetHotelByID", uint(1)).Return(expectedHotel, nil)

		app := fiber.New()
		app.Get("/hotel/:id", hotelHandler.GetHotelByID)

		req := httptest.NewRequest(http.MethodGet, "/hotel/1", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response utils.APIResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Hotel fetched successfully", response.Message)

		dataBytes, err := json.Marshal(response.Data)
		assert.NoError(t, err)

		var hotel model.HotelDetailResponse
		err = json.Unmarshal(dataBytes, &hotel)
		assert.NoError(t, err)

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		app := fiber.New()
		app.Get("/hotel/:id", hotelHandler.GetHotelByID)

		req := httptest.NewRequest(http.MethodGet, "/hotel/abc", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockService.On("GetHotelByID", uint(999)).Return(model.HotelDetailResponse{}, errors.New("not found"))

		app := fiber.New()
		app.Get("/hotel/:id", hotelHandler.GetHotelByID)

		req := httptest.NewRequest(http.MethodGet, "/hotel/999", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)

		var response utils.APIResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Hotel Not Found", response.Message)

		mockService.AssertExpectations(t)
	})
}
