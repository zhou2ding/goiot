package leetcode

func fibMemo(memo map[int]int, n int) int {
	if n <= 1 {
		return n
	}
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	memo[n] = fibMemo(memo, n-1) + fibMemo(memo, n-2)
	return memo[n]
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	fib1, fib2 := 1, 0
	for i := 2; i <= n; i++ {
		fibN := fib1 + fib2
		fib2 = fib1
		fib1 = fibN
	}
	return fib1
}
