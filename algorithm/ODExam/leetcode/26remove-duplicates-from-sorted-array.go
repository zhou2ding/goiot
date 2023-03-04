package leetcode

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return len(nums[:slow+1])
}
