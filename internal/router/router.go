package router

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/handler"
)

func SetupRoutes(app *fiber.App, hotelHandler *handler.HotelHandler, reviewHandler *handler.ReviewHandler) {
	app.Get("/hotels", hotelHandler.GetAllHotels)
	app.Get("/hotel/:id", hotelHandler.GetHotelByID)
	app.Post("/hotel", hotelHandler.CreateHotel)
	app.Put("/hotel/:id", hotelHandler.UpdateHotel)
	app.Delete("/hotel/:id", hotelHandler.DeleteHotel)

	app.Get("/hotel/:id/reviews", reviewHandler.GetReviewsByHotel)
	app.Post("/hotel/:id/reviews", reviewHandler.AddReview)
}
