package base_test

import (
	"testing"

	"github.com/coolestowl/leetcode/base"
)

func TestStack(t *testing.T) {
	s := base.NewStack[int]()

	for i := range make([]struct{}, 10) {
		s.Push(i)
	}

	if s.Top() != 9 {
		t.Errorf("top is not %d", 9)
	}

	for i := range make([]struct{}, 10) {
		expected := 9 - i
		if got := s.Pop(); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}

	ShouldPanic(t, func() {
		_ = s.Pop()
	})
	ShouldPanic(t, func() {
		_ = s.Top()
	})
}

func ShouldPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		}
	}()

	f()
}
