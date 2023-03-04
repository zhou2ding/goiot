package leetcode

import "strconv"

func openLock(deadends []string, target string) int {
	deads := make(map[string]bool)
	for _, d := range deadends {
		deads[d] = true
	}
	if deads["0000"] {
		return -1
	}

	queue := []string{"0000"}
	used2 := make(map[string]bool)
	step := 0
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			if deads[cur] {
				continue
			}
			for j := range cur {
				c := cur[j] - '0'
				var tmp string
				if c == 9 {
					tmp = cur[:j] + strconv.Itoa(int(c-9)) + cur[j+1:]
				} else {
					tmp = cur[:j] + strconv.Itoa(int(c+1)) + cur[j+1:]
				}
				if !used2[tmp] {
					used2[tmp] = true
					queue = append(queue, tmp)
				}
				if c == 0 {
					tmp = cur[:j] + strconv.Itoa(int(c+9)) + cur[j+1:]
				} else {
					tmp = cur[:j] + strconv.Itoa(int(c-1)) + cur[j+1:]
				}
				if !used2[tmp] {
					used2[tmp] = true
					queue = append(queue, tmp)
				}
			}
		}
		step++
	}
	return -1
}
