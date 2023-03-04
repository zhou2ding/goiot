package leetcode

import "math"

func maxArea(height []int) int {
	length := len(height)

	j := length - 1
	i := 0
	volume := (length - 1) * min(height[i], height[j])
	for i < j {
		if height[i] < height[j] {
			i++
			volume = cal(i, j, volume, height)
		} else {
			j--
			volume = cal(i, j, volume, height)
		}
	}

	return volume
}

func cal(i, j, maxVolume int, height []int) int {
	if min(height[i], height[j])*(j-i) > maxVolume {
		maxVolume = min(height[i], height[j]) * (j - i)
	}
	return maxVolume
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
