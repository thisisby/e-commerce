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
