/*
	双指针，左右各一个，最大容量设为初始容积
	容积S(i,j) = (j - i) * min(i,j)，底边是(j-i)，侧边是min(i,j)
	每次向内移动时，若移动较长板，则min(i,j)不变或缩小，且(j-i)必缩小，因此容积必缩小；
				若移动较短板，则min(i,j)不变、缩小、增大均有可能，虽然(j-i)必缩小，但容积有增大的可能；
				因此每次只移动较短板
*/
package main

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
