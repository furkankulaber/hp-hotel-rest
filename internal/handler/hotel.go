package handler

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/service"
	"strconv"
)

type HotelHandler struct {
	service service.HotelService
}

func NewHotelHandler(service service.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func (h *HotelHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/hotels", h.GetAllHotels)
	app.Get("/hotels/:id", h.GetHotelByID)
	app.Post("/hotels", h.CreateHotel)
	app.Put("/hotels/:id", h.UpdateHotel)
	app.Delete("/hotels/:id", h.DeleteHotel)
}

func (h *HotelHandler) GetAllHotels(c *fiber.Ctx) error {
	hotels, err := h.service.GetAllHotels()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) GetHotelByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Hotel ID"})
	}
	hotel, err := h.service.GetHotelByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Hotel Not Found"})
	}
	return c.JSON(hotel)
}

func (h *HotelHandler) CreateHotel(c *fiber.Ctx) error {
	var hotel model.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}
	createdHotel, err := h.service.CreateHotel(hotel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusCreated).JSON(createdHotel)
}

func (h *HotelHandler) UpdateHotel(c *fiber.Ctx) error {
	var hotel model.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	hotel.ID = uint(id)

	updatedHotel, err := h.service.UpdateHotel(hotel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(updatedHotel)
}

func (h *HotelHandler) DeleteHotel(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Hotel ID"})
	}
	err = h.service.DeleteHotel(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Hotel Not Found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
