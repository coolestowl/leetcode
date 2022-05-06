package sword2offer

import (
	"github.com/coolestowl/leetcode/base"
)

type StackToQueue struct {
	s1 *base.Stack
	s2 *base.Stack
}

func NewStackToQueue() *StackToQueue {
	return &StackToQueue{
		s1: base.NewStack(),
		s2: base.NewStack(),
	}
}

func (s *StackToQueue) AppendTail(elem base.ElemType) {
	s.s1.Push(elem)
}

func (s *StackToQueue) DeleteHead() base.ElemType {
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

type minItem struct {
	current base.ElemType
	min     base.ElemType
}

type MinStack struct {
	inner  []minItem
	topPtr int
}

func NewMinStack() *MinStack {
	return &MinStack{
		inner:  make([]minItem, 0),
		topPtr: -1,
	}
}

func (s *MinStack) Push(elem base.ElemType) {
	min := elem
	if len(s.inner) > 0 {
		topItem := s.inner[s.topPtr]
		if topItem.min < min {
			min = topItem.min
		}
	}

	s.inner = append(s.inner, minItem{
		current: elem,
		min:     min,
	})
	s.topPtr++
}

func (s *MinStack) Pop() base.ElemType {
	if len(s.inner) > 0 {
		item := s.inner[s.topPtr]
		s.inner = s.inner[:s.topPtr]
		s.topPtr--
		return item.current
	}

	return -1
}

func (s *MinStack) Top() base.ElemType {
	if len(s.inner) > 0 {
		item := s.inner[s.topPtr]
		return item.current
	}

	return -1
}

func (s *MinStack) Min() base.ElemType {
	if len(s.inner) > 0 {
		item := s.inner[s.topPtr]
		return item.min
	}

	return -1
}
