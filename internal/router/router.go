package router

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/handler"
)

func SetupRoutes(app *fiber.App, hotelHandler *handler.HotelHandler, reviewHandler *handler.ReviewHandler) {
	app.Get("/hotels", hotelHandler.GetAllHotels)
	app.Get("/hotel/:id", hotelHandler.GetHotelByID)

	app.Get("/hotel/:id/reviews", reviewHandler.GetReviewsByHotel)
	app.Post("/hotel/:id/reviews", reviewHandler.AddReview)
}
