package leetcode

var result [][]int
var path []int
var used []bool

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	path = make([]int, 0)
	used = make([]bool, len(nums))
	dfs(nums)
	return result
}

func dfs(nums []int) {
	if len(path) == len(nums) {
		result = append(result, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
