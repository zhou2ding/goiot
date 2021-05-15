package utils

import "strconv"

func Str2Int(s string) int {
	res, _ := strconv.ParseInt(s, 10, 0)
	return int(res)
}
