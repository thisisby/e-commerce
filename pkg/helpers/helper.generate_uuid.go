package helpers

import "github.com/google/uuid"

func GenerateUUID() string {
	uid := uuid.New()

	return uid.String()
}
