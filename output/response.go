package output

import (
	"ReMotion-C7/app/dto/response"

	"github.com/gofiber/fiber/v2"
)

func GetOutput(c *fiber.Ctx, status string, fiberStatus int, message string, data interface{}) error {
	response := response.ApiResponse{
		Status:  status,
		Message: message,
	}

	if data != nil {
		response.Data = data
	}

	return c.Status(fiberStatus).JSON(response)
}
