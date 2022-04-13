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
