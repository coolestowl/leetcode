package sword2offer

func ReverseLeftWords(s string, n int) string {
	if len(s) <= 1 {
		return s
	}
	n = n % len(s)
	return s[n:] + s[:n]
}
