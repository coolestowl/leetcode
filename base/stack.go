package base

type Stack[T any] struct {
	inner  []T
	topPtr int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		inner:  make([]T, 0),
		topPtr: -1,
	}
}

func (s *Stack[T]) Push(elem T) {
	s.inner = append(s.inner, elem)
	s.topPtr++
}

func (s *Stack[T]) Pop() T {
	if s.Len() > 0 {
		elem := s.inner[s.topPtr]
		s.inner = s.inner[:s.topPtr]
		s.topPtr--
		return elem
	}

	panic("empty stack")
}

func (s *Stack[T]) Len() int {
	return len(s.inner)
}

func (s *Stack[T]) Top() T {
	if s.Len() > 0 {
		return s.inner[s.topPtr]
	}

	panic("empty stack")
}
