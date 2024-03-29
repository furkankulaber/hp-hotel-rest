package handler

import (
	"github.com/gofiber/fiber/v2"
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
