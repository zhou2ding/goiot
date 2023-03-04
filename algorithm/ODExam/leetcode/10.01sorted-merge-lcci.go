package leetcode

func merge(A []int, m int, B []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(A, B)
		return
	}

	i1, i2, i := 0, 0, 0
	resultTmp := make([]int, m+n)
	for i1 < m && i2 < n {
		if A[i1] < B[i2] {
			resultTmp[i] = A[i1]
			i1++
		} else {
			resultTmp[i] = B[i2]
			i2++
		}
		i++
	}
	if i1 == m {
		for j := i2; j < n; j++ {
			resultTmp[i] = B[j]
			i++
		}
	}
	if i2 == n {
		for j := i1; j < m; j++ {
			resultTmp[i] = A[j]
			i++
		}
	}
	copy(A, resultTmp)
}
