package leetcode

var (
	res4  [][]int
	path4 []int
	sum   int
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 || (len(candidates) == 1 && candidates[0] != target) {
		return nil
	}

	res4 = make([][]int, 0)
	path4 = make([]int, 0)
	dfs4(candidates, 0, target)
	return res4
}

func dfs4(candidates []int, start, target int) {
	if sum == target {
		res4 = append(res4, append([]int{}, path4...))
		return
	}
	if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		sum += candidates[i]
		path4 = append(path4, candidates[i])
		dfs4(candidates, i, target)
		sum -= path4[len(path4)-1]
		path4 = path4[:len(path4)-1]
	}
}
