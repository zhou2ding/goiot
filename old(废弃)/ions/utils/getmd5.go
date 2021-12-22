package utils

import "crypto/md5"

func GetMd5File(s string) string {
	tmp := md5.Sum([]byte(s))
	ret := string(tmp[:])
	return ret
}
