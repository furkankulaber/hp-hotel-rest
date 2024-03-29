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

// AddReview godoc
// @Summary Add a new review
// @Description Add a new review for a specific hotel
// @Tags reviews
// @Accept json
// @Produce json
// @Param hotelID path integer true "Hotel ID"
// @Param request body model.CreateReviewRequest true "Review request body"
// @Success 201 {object} utils.APIResponse{data=model.ReviewResponse} "Review added successfully"
// @Failure 400 {object} utils.APIResponse "Bad Request"
// @Failure 500 {object} utils.APIResponse "Internal Server Error"
// @Router /hotel/{hotelID}/reviews [post]
func (h *ReviewHandler) AddReview(c *fiber.Ctx) error {
	var request model.CreateReviewRequest
	if err := c.BodyParser(&request); err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Bad Request", nil)
	}

	var hotel model.Hotel
	id, err := strconv.ParseUint(c.Params("hotelID"), 10, 32)
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

// GetReviewsByHotel godoc
// @Summary Get reviews by hotel ID
// @Description Get reviews for a specific hotel
// @Tags reviews
// @Accept json
// @Produce json
// @Param hotelID path integer true "Hotel ID"
// @Success 200 {array} utils.APIResponse{data=[]model.ReviewResponse} "Reviews fetched successfully"
// @Failure 400 {object} utils.APIResponse "Invalid hotel ID"
// @Failure 500 {object} utils.APIResponse "Internal Server Error"
// @Router /hotel/{hotelID}/reviews [get]
func (h *ReviewHandler) GetReviewsByHotel(c *fiber.Ctx) error {
	hotelID, err := c.ParamsInt("hotelID")
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Invalid hotel ID", nil)
	}

	reviews, err := h.service.GetReviewsByHotelID(uint(hotelID))
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusInternalServerError, "Could not retrieve reviews", nil)
	}

	return utils.RespondJSON(c, fiber.StatusOK, "Reviews fetched successfully", reviews)
}

// UpdateReview godoc
// @Summary Update a review
// @Description Update an existing review
// @Tags reviews
// @Accept json
// @Produce json
// @Param reviewID path integer true "Review ID"
// @Param request body model.UpdateReviewRequest true "Review request body"
// @Success 200 {object} utils.APIResponse{data=model.ReviewResponse} "Review updated successfully"
// @Failure 400 {object} utils.APIResponse "Bad Request"
// @Failure 500 {object} utils.APIResponse "Internal Server Error"
// @Router /hotel/reviews/{reviewID} [put]
func (h *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	var req model.UpdateReviewRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})
	}

	reviewID, err := strconv.ParseUint(c.Params("reviewID"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid review ID"})
	}

	updatedReview, err := h.service.UpdateReview(uint(reviewID), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(updatedReview)
}
