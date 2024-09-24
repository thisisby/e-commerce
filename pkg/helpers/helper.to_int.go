package helpers

import "strconv"

func ToInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}

	return defaultValue
}

func ToFloat64(value string, defaultValue float64) float64 {
	if value == "" {
		return defaultValue
	}
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return floatValue
	}

	return defaultValue
}
