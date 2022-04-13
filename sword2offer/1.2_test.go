package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/sword2offer"
)

func TestMinStack(t *testing.T) {
	s := sword2offer.NewMinStack[int]()

	for i := range make([]struct{}, 10) {
		s.Push(i)

		if got := s.Min(); got != 0 {
			t.Errorf("expected %d, got %d", 0, got)
		}
	}

	if got := s.Top(); got != 9 {
		t.Errorf("top is not %d", 9)
	}

	for i := range make([]struct{}, 10) {
		expected := 9 - i
		if got := s.Pop(); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}

	for _, f := range []func() int{
		s.Pop, s.Top, s.Min,
	} {
		if got := f(); got != -1 {
			t.Errorf("expected %d, got %d", -1, got)
		}
	}

	for i := range make([]struct{}, 10) {
		elem := 9 - i

		s.Push(elem)

		if got := s.Min(); got != elem {
			t.Errorf("expected %d, got %d", elem, got)
		}
	}
}
