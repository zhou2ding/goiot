package utils

import uuid "github.com/google/uuid"

func GetUUIDFull() string {
	return uuid.New().String()
}
