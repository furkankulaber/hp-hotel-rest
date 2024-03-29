package handler

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/service"
	"strconv"
)

type ReviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) AddReview(c *fiber.Ctx) error {
	var request model.CreateReviewRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var hotel model.Hotel

	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	hotel.ID = uint(id)

	newReview, err := h.service.AddReview(&request, &hotel)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newReview)
}

func (h *ReviewHandler) GetReviewsByHotel(c *fiber.Ctx) error {
	hotelID, err := c.ParamsInt("idg")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid hotel ID"})
	}

	reviews, err := h.service.GetReviewsByHotelID(uint(hotelID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve reviews"})
	}

	return c.JSON(reviews)
}
