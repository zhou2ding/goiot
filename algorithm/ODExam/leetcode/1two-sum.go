package leetcode

func twoSum(nums []int, target int) []int {
	m := getSubVal(nums, target)
	for i, v := range nums {
		if idx, ok := m[v]; ok {
			if i != idx {
				return []int{i, idx}
			}
		}
	}
	return nil
}

func getSubVal(nums []int, target int) map[int]int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[target-v] = i
	}
	return m
}
