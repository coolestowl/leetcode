package sword2offer_test

import (
	"math/rand"
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

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
