package utils

import "github.com/gofiber/fiber/v2"

func RespondJSON(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode >= 200 && statusCode < 300,
		"message": message,
		"data":    data,
	})
}
