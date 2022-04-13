package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/base"
	"github.com/coolestowl/leetcode/sword2offer"
)

func TestReverse(t *testing.T) {
	head := base.NewLinkList(-1)

	for i := range make([]struct{}, 10) {
		head.AppendTail(i)
		head.InsertNext(i)
	}

	head = sword2offer.Reverse(head)

	expected := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		-1,
	}
	reversed := make([]int, 0, len(expected))
	for ptr := head; ptr != nil; ptr = ptr.Next {
		reversed = append(reversed, ptr.Val)
	}

	if !sliceEqual(expected, reversed) {
		t.Errorf("expected %v, got %v", expected, reversed)
	}

	if got := sword2offer.Reverse[int](nil); got != nil {
		t.Errorf("expected nil")
	}
}
