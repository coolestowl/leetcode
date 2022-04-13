package sword2offer

import "github.com/coolestowl/leetcode/base"

type minItem[T Elem] struct {
	current T
	min     T
}

type MinStack[T Elem] struct {
	inner *base.Stack[minItem[T]]
}

func NewMinStack[T Elem]() *MinStack[T] {
	return &MinStack[T]{
		inner: base.NewStack[minItem[T]](),
	}
}

func (s *MinStack[T]) Push(elem T) {
	min := elem
	if s.inner.Len() > 0 {
		topItem := s.inner.Top()
		if topItem.min < min {
			min = topItem.min
		}
	}

	s.inner.Push(minItem[T]{
		current: elem,
		min:     min,
	})
}

func (s *MinStack[T]) Pop() T {
	if s.inner.Len() > 0 {
		item := s.inner.Pop()
		return item.current
	}

	return -1
}

func (s *MinStack[T]) Top() T {
	if s.inner.Len() > 0 {
		item := s.inner.Top()
		return item.current
	}

	return -1
}

func (s *MinStack[T]) Min() T {
	if s.inner.Len() > 0 {
		item := s.inner.Top()
		return item.min
	}

	return -1
}
