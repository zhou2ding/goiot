package utils

import uuid "github.com/satori/go.uuid"

func GetUUIDFull() string {
	return uuid.NewV4().String()
}
