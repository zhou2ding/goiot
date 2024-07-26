package utils

import (
	"net"
	"time"
)

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

func SliceToMap[T comparable](slice []T) map[T]bool {
	result := make(map[T]bool)
	for _, elem := range slice {
		result[elem] = true
	}
	return result
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
