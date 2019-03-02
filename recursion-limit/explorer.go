package main

func fib(n int, memo []int64) int64 {
	var result int64
	if memo[n] != 0 {
		print("hit memoize ")
		print(n)
		println()
		return memo[n]
	}
	if n <= 2 {
		result = 1
	} else {
		result = fib(n-1, memo) + fib(n-2, memo)
	}
	memo[n] = result
	return result
}

func main() {
	n := 8999999 // goroutine stack exceeds 1000000000-byte limit :))
	memo := make([]int64, n+1)
	print(fib(n, memo))
}
