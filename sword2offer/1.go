package sword2offer

import (
	"github.com/coolestowl/leetcode/base"
)

type Elem interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 // and more, maybe any
}

type StackToQueue[T Elem] struct {
	s1 *base.Stack[T]
	s2 *base.Stack[T]
}

func NewStackToQueue[T Elem]() *StackToQueue[T] {
	return &StackToQueue[T]{
		s1: base.NewStack[T](),
		s2: base.NewStack[T](),
	}
}

func (s *StackToQueue[T]) AppendTail(elem T) {
	s.s1.Push(elem)
}

func (s *StackToQueue[T]) DeleteHead() T {
	if s.s2.Len() > 0 {
		return s.s2.Pop()
	}

	for s.s1.Len() > 0 {
		s.s2.Push(s.s1.Pop())
	}

	if s.s2.Len() > 0 {
		return s.s2.Pop()
	}

	return -1
}

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