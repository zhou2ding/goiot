package utils

import "time"

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func CalCosts(start int64) int64 {
	return time.Now().Sub(time.Unix(start, 0)).Milliseconds()
}
