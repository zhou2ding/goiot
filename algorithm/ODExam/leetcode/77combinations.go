package leetcode

var res3 [][]int
var path3 []int

func combine(n int, k int) [][]int {
	res3 = make([][]int, 0)
	path3 = make([]int, 0)
	dfs3(n, k, 1)
	return res3
}

func dfs3(n, k, start int) {
	if len(path3) == k {
		res3 = append(res3, append([]int{}, path3...))
	}
	for i := start; i <= n; i++ {
		path3 = append(path3, i)
		dfs3(n, k, i+1)
		path3 = path3[:len(path3)-1]
	}
}
