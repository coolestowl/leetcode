package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/base"
	"github.com/coolestowl/leetcode/sword2offer"
)

func TestReversePrint(t *testing.T) {
	head := base.NewLinkList(-1)

	for i := range make([]struct{}, 10) {
		head.AppendTail(i)
		head.InsertNext(i)
	}

	expected := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		-1,
	}
	reversed := sword2offer.ReversePrint(head)
	if !sliceEqual(expected, reversed) {
		t.Errorf("expected %v, got %v", expected, reversed)
	}
}

func sliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, elem := range a {
		if elem != b[idx] {
			return false
		}
	}
	return true
}
