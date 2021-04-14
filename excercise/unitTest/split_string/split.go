package split_string

import "strings"

func Split(s, sep string) []string {
	var ret = make([]string, 0, strings.Count(s, sep)+1)
	for {
		index := strings.Index(s, sep)
		if index < 0 {
			break
		}
		ret = append(ret, s[:index])
		s = s[index+len(sep):]
	}
	ret = append(ret, s)
	return ret
}
