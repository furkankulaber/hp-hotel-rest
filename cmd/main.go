// cmd/app/main.go

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "hp-hotel-rest/docs"
	"hp-hotel-rest/internal/handler"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/repository"
	"hp-hotel-rest/internal/router"
	"hp-hotel-rest/internal/seed"
	"hp-hotel-rest/internal/service"
	"hp-hotel-rest/pkg/config"
)

// @title HotelPro Hotel REST API
// @description This is a sample API for managing hotels and reviews.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @SecurityDefinitions.apiKey Bearer
// @In header
// @Name Authorization
// @Description "Please using <b>Bearer: JWT</b>"
// @Value Bearer:
// @Type apiKey

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.Hotel{}, &model.Location{}, &model.Amenity{}, &model.Photo{}, &model.Review{}, &model.Room{}, &model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	seed.Load(db)

	hotelRepo := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepo)
	hotelHandler := handler.NewHotelHandler(hotelService)

	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router.SetupRoutes(app, hotelHandler, reviewHandler, authHandler)

	log.Fatal(app.Listen(":8080"))
}
