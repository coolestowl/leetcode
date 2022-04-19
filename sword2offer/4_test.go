package sword2offer_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

func TestFindRepeatNumber(t *testing.T) {
	cases := []struct {
		Nums     []int
		Expected int
	}{{
		Nums:     []int{1, 2, 3, 4, 4, 5, 6, 6, 7},
		Expected: 4,
	}, {
		Nums:     []int{7, 6, 5, 4, 3, 2},
		Expected: -1,
	}}

	for _, c := range cases {
		if got := sword2offer.FindRepeatNumber(c.Nums); got != c.Expected {
			t.Errorf("FindRepeatNumber(%v) == %v, want %v", c.Nums, got, c.Expected)
		}
	}
}

func TestSearch(t *testing.T) {
	cases := []struct {
		Nums     []int
		Target   int
		Expected int
	}{{
		Nums:     []int{2, 3, 5, 6, 7, 8, 8, 9},
		Target:   8,
		Expected: 2,
	}, {
		Nums:     []int{2, 3, 4, 6, 7, 8, 8, 9},
		Target:   5,
		Expected: 0,
	}}

	for _, c := range cases {
		if got := sword2offer.Search(c.Nums, c.Target); got != c.Expected {
			t.Errorf("Search(%v, %v) == %v, want %v", c.Nums, c.Target, got, c.Expected)
		}
	}
}

func FuzzBinSearch(f *testing.F) {
	f.Fuzz(func(t *testing.T, nums []int, target, last int) {
	})
}

func TestBinSearch(t *testing.T) {
	indexOf := func(elems []int, target int) int {
		for idx, elem := range elems {
			if elem == target {
				return idx
			}
		}
		return -1
	}
	lastIndexOf := func(elems []int, target int) int {
		for i := len(elems) - 1; i >= 0; i-- {
			if elems[i] == target {
				return i
			}
		}
		return -1
	}

	N := 10000
	for range make([]struct{}, N) {
		n := rand.Intn(N)

		nums := make([]int, 0, n)
		for range make([]struct{}, n) {
			nums = append(nums, rand.Intn(N))
		}
		sort.Ints(nums)

		target := rand.Intn(N)
		last := rand.Intn(2)

		var expected int
		if last == 0 {
			expected = indexOf(nums, target)
		} else {
			expected = lastIndexOf(nums, target)
		}

		if got := sword2offer.BinSearch(nums, target, last); got != expected {
			t.Errorf("BinSearch(%v, %v, %v) == %v, want %v", nums, target, last, got, expected)
		}
	}
}
