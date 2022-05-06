package sword2offer_test

import (
	"testing"

	"github.com/coolestowl/leetcode/base"
	"github.com/coolestowl/leetcode/sword2offer"
)

func TestStackToQueue(t *testing.T) {
	s := sword2offer.NewStackToQueue()

	for i := range make([]struct{}, 10) {
		s.AppendTail(base.ElemType(i))
	}

	for i := range make([]struct{}, 10) {
		if got := s.DeleteHead(); got != base.ElemType(i) {
			t.Errorf("expected %d got %d", i, got)
		}
	}

	for i := range make([]struct{}, 10) {
		s.AppendTail(base.ElemType(i))
	}

	for i := range make([]struct{}, 5) {
		if got := s.DeleteHead(); got != base.ElemType(i) {
			t.Errorf("expected %d got %d", i, got)
		}
	}

	for i := range make([]struct{}, 10) {
		s.AppendTail(base.ElemType(i))
	}

	for _, i := range []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if got := s.DeleteHead(); got != base.ElemType(i) {
			t.Errorf("expected %d got %d", i, got)
		}
	}

	for range make([]struct{}, 5) {
		if got := s.DeleteHead(); got != -1 {
			t.Errorf("expected %d got %d", -1, got)
		}
	}
}

func TestMinStack(t *testing.T) {
	s := sword2offer.NewMinStack()

	for i := range make([]struct{}, 10) {
		s.Push(base.ElemType(i))

		if got := s.Min(); got != 0 {
			t.Errorf("expected %d, got %d", 0, got)
		}
	}

	if got := s.Top(); got != 9 {
		t.Errorf("top is not %d", 9)
	}

	for i := range make([]struct{}, 10) {
		expected := base.ElemType(9 - i)
		if got := s.Pop(); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}

	for _, f := range []func() base.ElemType{
		s.Pop, s.Top, s.Min,
	} {
		if got := f(); got != -1 {
			t.Errorf("expected %d, got %d", -1, got)
		}
	}

	for i := range make([]struct{}, 10) {
		elem := base.ElemType(9 - i)

		s.Push(elem)

		if got := s.Min(); got != elem {
			t.Errorf("expected %d, got %d", elem, got)
		}
	}
}
