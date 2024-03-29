package utils

import "github.com/gofiber/fiber/v2"

func RespondJSON(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	status := "failed"
	if statusCode >= 200 && statusCode < 300 {
		status = "success"
	}
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
