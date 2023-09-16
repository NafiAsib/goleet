package leetcode

var dp = make(map[int]int)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if val, ok := dp[n]; ok { //comma-ok idiom
		return val
	}

	dp[n] = fib(n-1) + fib(n-2)

	return dp[n]
}
