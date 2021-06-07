// 11. 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/

package main

import (
	"math"
	"sort"
)

func maxArea(height []int) int {
	var (
		ret int
		maxVolume int
	)
	length := len(height)

	if justRet1(height) || justRet2(height) {
		ret = (length - 1) * height[0]
		return ret
	}

	for i := 0; i < length - 1; i++ {
		for j := 1; j < length; j++ {
			tmp := int(math.Min(float64(height[i]),float64(height[j])))
			if (j - i) * tmp >= maxVolume {
				maxVolume = (j - i) * tmp
			}
		}
	}

	return maxVolume
}

func justRet2(height []int) bool {
	length := len(height)
	tmp := make([]int,length,length)
	copy(tmp,height)
	sort.Ints(tmp)
	max := tmp[length - 1]
	if height[0] == height[length - 1] && height[0]  == max {
		return true
	}
	return false
}

func justRet1(height []int) bool {
	for i := 1; i < len(height);i++ {
		if height[i] != height[i-1] {
			return false
		}
	}
	return true
}