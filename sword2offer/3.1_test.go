package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

func TestReplaceSpace(t *testing.T) {
	cases := []struct {
		Arg      string
		Expected string
	}{{
		Arg:      "hello world",
		Expected: "hello%20world",
	}, {
		Arg:      "   ",
		Expected: "%20%20%20",
	}, {
		Arg:      "",
		Expected: "",
	}}

	for _, c := range cases {
		if got := sword2offer.ReplaceSpace(c.Arg); got != c.Expected {
			t.Errorf("ReplaceSpace(%q) == %q, want %q", c.Arg, got, c.Expected)
		}
	}
}
