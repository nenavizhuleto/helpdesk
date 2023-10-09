package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleError(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	err = c.Status(code).JSON(ErrorDTO{Status: code, Message: e.Message})
	if err != nil {
		// In case the SendFile fails
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorDTO{Status: code, Message: "Internal server error"})
	}

	// Return from handler
	return nil
}
