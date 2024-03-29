package router

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/handler"
)

func SetupRoutes(app *fiber.App, hotelHandler *handler.HotelHandler) {
	// Hotel ile ilgili route'lar
	app.Get("/hotels", hotelHandler.GetAllHotels)
	app.Get("/hotels/:id", hotelHandler.GetHotelByID)
	app.Post("/hotels", hotelHandler.CreateHotel)
	app.Put("/hotels/:id", hotelHandler.UpdateHotel)
	app.Delete("/hotels/:id", hotelHandler.DeleteHotel)

	// Daha fazla route buraya eklenebilir
}
