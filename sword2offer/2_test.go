package sword2offer_test

import (
	"math/rand"
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

func TestCopyRandomList(t *testing.T) {
	N := 10

	items := make([]sword2offer.NodeItem[int], 0, N)

	for i := range make([]struct{}, N) {
		items = append(items, sword2offer.NodeItem[int]{
			Val: i,
			RandomIdx: func() int {
				randIdx := rand.Intn(N + 1)
				if randIdx == N {
					randIdx = -1
				}
				return randIdx
			}(),
		})
	}

	head := sword2offer.NodeItemsToLinkList(items)

	copied := sword2offer.CopyRandomList(head)

	for ptr := head; ptr != nil; ptr, copied = ptr.Next, copied.Next {
		if copied == nil {
			t.Errorf("copied is nil")
		}

		if ptr.Val != copied.Val {
			t.Errorf("copied.Val is %d, but ptr.Val is %d", copied.Val, ptr.Val)
		}

		if ptr.Random != nil {
			if copied.Random == nil {
				t.Errorf("copied.Random is nil")
			}

			if ptr.Random.Val != copied.Random.Val {
				t.Errorf("copied.Random.Val is %d, but ptr.Random.Val is %d", copied.Random.Val, ptr.Random.Val)
			}
		}
	}

	copied = sword2offer.CopyRandomList[int](nil)
	if copied != nil {
		t.Errorf("copied is not nil")
	}
}
