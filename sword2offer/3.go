package sword2offer

import "strings"

func ReplaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func ReverseLeftWords(s string, n int) string {
	if len(s) <= 1 {
		return s
	}
	n = n % len(s)
	return s[n:] + s[:n]
}
