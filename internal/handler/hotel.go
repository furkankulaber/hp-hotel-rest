package handler

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/service"
	"hp-hotel-rest/pkg/utils"
	"strconv"
)

type HotelHandler struct {
	service service.HotelService
}

func NewHotelHandler(service service.HotelService) *HotelHandler {
	return &HotelHandler{service: service}
}

func (h *HotelHandler) GetAllHotels(c *fiber.Ctx) error {
	hotels, err := h.service.GetAllHotels()
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Internal Server Error", nil)
	}
	return utils.RespondJSON(c, fiber.StatusOK, "Hotels fetched successfully", hotels)
}

func (h *HotelHandler) GetHotelByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Hotel ID"})
	}
	hotel, err := h.service.GetHotelByID(uint(id))
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusNotFound, "Hotel Not Found", nil)
	}
	return utils.RespondJSON(c, fiber.StatusOK, "Hotel fetched successfully", hotel)
}

func (h *HotelHandler) CreateHotel(c *fiber.Ctx) error {
	var hotel model.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Bad Request", nil)
	}
	createdHotel, err := h.service.CreateHotel(hotel)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}
	return utils.RespondJSON(c, fiber.StatusCreated, "Hotel created successfully", createdHotel)
}

func (h *HotelHandler) UpdateHotel(c *fiber.Ctx) error {
	var hotel model.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Bad Request", nil)
	}
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	hotel.ID = uint(id)

	updatedHotel, err := h.service.UpdateHotel(hotel)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}
	return utils.RespondJSON(c, fiber.StatusOK, "Hotel updated successfully", updatedHotel)
}

func (h *HotelHandler) DeleteHotel(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Invalid Hotel ID", nil)
	}
	err = h.service.DeleteHotel(uint(id))
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusNotFound, "Hotel Not Found", nil)
	}
	return utils.RespondJSON(c, fiber.StatusNoContent, "Hotel deleted successfully", nil)
}
