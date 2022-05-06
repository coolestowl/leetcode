package base

type Stack struct {
	inner  []ElemType
	topPtr int
}

func NewStack() *Stack {
	return &Stack{
		inner:  make([]ElemType, 0),
		topPtr: -1,
	}
}

func (s *Stack) Push(elem ElemType) {
	s.inner = append(s.inner, elem)
	s.topPtr++
}

func (s *Stack) Pop() ElemType {
	if s.Len() > 0 {
		elem := s.inner[s.topPtr]
		s.inner = s.inner[:s.topPtr]
		s.topPtr--
		return elem
	}

	panic("empty stack")
}

func (s *Stack) Len() int {
	return len(s.inner)
}

func (s *Stack) Top() ElemType {
	if s.Len() > 0 {
		return s.inner[s.topPtr]
	}

	panic("empty stack")
}
