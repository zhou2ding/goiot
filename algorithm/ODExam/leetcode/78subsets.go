package leetcode

var res2 [][]int
var path2 []int

func subsets(nums []int) [][]int {
	res2 = make([][]int, 0)
	path2 = make([]int, 0)
	dfs2(nums, 0)
	return res2
}

func dfs2(nums []int, start int) {
	res2 = append(res2, append([]int{}, path2...))
	for i := start; i < len(nums); i++ {
		path2 = append(path2, nums[i])
		dfs2(nums, i+1)
		path2 = path2[:len(path2)-1]
	}
}
