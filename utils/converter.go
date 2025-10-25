package utils

import (
	"strconv"
)

func ConvertToNum(idStr string) (int, error) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
