package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

func TestStackToQueue(t *testing.T) {
	s := sword2offer.NewStackToQueue[int]()

	for i := range make([]struct{}, 10) {
		s.AppendTail(i)
	}

	for i := range make([]struct{}, 10) {
		if got := s.DeleteHead(); got != i {
			t.Errorf("expected %d got %d", i, got)
		}
	}

	for i := range make([]struct{}, 10) {
		s.AppendTail(i)
	}

	for i := range make([]struct{}, 5) {
		if got := s.DeleteHead(); got != i {
			t.Errorf("expected %d got %d", i, got)
		}
	}

	for i := range make([]struct{}, 10) {
		s.AppendTail(i)
	}

	for _, i := range []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if got := s.DeleteHead(); got != i {
			t.Errorf("expected %d got %d", i, got)
		}
	}

	for range make([]struct{}, 5) {
		if got := s.DeleteHead(); got != -1 {
			t.Errorf("expected %d got %d", -1, got)
		}
	}
}
