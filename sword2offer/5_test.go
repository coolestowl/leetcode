package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

func TestFindNumberIn2DArray(t *testing.T) {
	array := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 20

	if got := sword2offer.FindNumberIn2DArray(array, target); got {
		t.Errorf("???")
	}
}

func TestHHH(t *testing.T) {
	mid := (2 + 3) / 2
	t.Error(mid)
}
