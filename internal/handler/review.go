package handler

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/service"
	"hp-hotel-rest/pkg/utils"
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
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Bad Request", nil)
	}

	var hotel model.Hotel
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Invalid hotel ID", nil)
	}
	hotel.ID = uint(id)

	newReview, err := h.service.AddReview(&request, &hotel)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}

	return utils.RespondJSON(c, fiber.StatusCreated, "Review added successfully", newReview)
}

func (h *ReviewHandler) GetReviewsByHotel(c *fiber.Ctx) error {
	hotelID, err := c.ParamsInt("id")
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Invalid hotel ID", nil)
	}

	reviews, err := h.service.GetReviewsByHotelID(uint(hotelID))
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusInternalServerError, "Could not retrieve reviews", nil)
	}

	return utils.RespondJSON(c, fiber.StatusOK, "Reviews fetched successfully", reviews)
}

func (h *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	var req model.UpdateReviewRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	reviewID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid review ID"})
	}

	updatedReview, err := h.service.UpdateReview(uint(reviewID), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(updatedReview)
}
