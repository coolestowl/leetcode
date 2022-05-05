package sword2offer

func SumNums(n int) int {
	_ = n > 0 && func() bool { n = n + SumNums(n-1); return true }()
	return n
}
