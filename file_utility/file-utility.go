package file_utility

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) { // To make it be publicly available for access
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("Failed to read file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored value")
	}
	return value, nil
}
