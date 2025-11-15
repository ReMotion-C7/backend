package utils

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PointNumber(num int) *uint {
	v := uint(num)
	return &v
}

func ConvertToNum(c *fiber.Ctx, idStr string) (int, error) {

	id := c.Params(idStr)

	fmt.Println(id)

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return idNum, nil

}

func ConvertToNum2Var(c *fiber.Ctx, id1Str string, id2Str string) (int, int, error) {

	id1Num, err := ConvertToNum(c, id1Str)
	if err != nil {
		return 0, 0, err
	}

	id2Num, err := ConvertToNum(c, id2Str)
	if err != nil {
		return 0, 0, err
	}

	return id1Num, id2Num, nil

}
