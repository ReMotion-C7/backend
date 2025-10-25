package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ConvertToNum(c *fiber.Ctx, idStr string) (int, error) {

	id := c.Params(idStr)

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return idNum, nil
	
}
