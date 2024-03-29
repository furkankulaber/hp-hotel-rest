package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/fiber-swagger"
	"hp-hotel-rest/internal/handler"
	middlewares "hp-hotel-rest/internal/middleware"
)

func SetupRoutes(app *fiber.App, hotelHandler *handler.HotelHandler, reviewHandler *handler.ReviewHandler, authHandler *handler.AuthHandler) {

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Get("/hotels", hotelHandler.GetAllHotels)
	app.Get("/hotel/:id", hotelHandler.GetHotelByID)

	app.Get("/hotel/:hotelID/reviews", reviewHandler.GetReviewsByHotel)
	app.Post("/hotel/:hotelID/reviews", reviewHandler.AddReview)
	app.Put("/hotel/reviews/:reviewID", reviewHandler.UpdateReview)

	app.Post("/auth/register", authHandler.Register)
	app.Post("/auth/login", authHandler.Login)
	app.Get("/auth/protected", middlewares.NewAuthMiddleware(), authHandler.ProtectedRoute)
}
