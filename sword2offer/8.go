package sword2offer

func Fib(n int) int {
	// 0, 1, 1, 2, 3, 5, 8, 13
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}

func NumWays(n int) int {
	// 1, 2, 3, 5, 8, ...
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}
