package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

func TestReverseLeftWords(t *testing.T) {
	cases := []struct {
		Arg      string
		N        int
		Expected string
	}{{
		Arg:      "hello",
		N:        0,
		Expected: "hello",
	}, {
		Arg:      "hello",
		N:        2,
		Expected: "llohe",
	}, {
		Arg:      "hello",
		N:        233,
		Expected: "lohel",
	}, {
		Arg:      "a",
		N:        2333,
		Expected: "a",
	}}

	for _, c := range cases {
		if got := sword2offer.ReverseLeftWords(c.Arg, c.N); got != c.Expected {
			t.Errorf("ReverseLeftWords(%q, %d) == %q, want %q", c.Arg, c.N, got, c.Expected)
		}
	}
}
