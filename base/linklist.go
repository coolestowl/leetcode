package base

type LinkListNode[T any] struct {
	Val  T
	Next *LinkListNode[T]
}

func NewLinkList[T any](val T) *LinkListNode[T] {
	return &LinkListNode[T]{
		Val:  val,
		Next: nil,
	}
}

func (l *LinkListNode[T]) AppendTail(val T) {
	if l == nil {
		return
	}

	var ptr *LinkListNode[T]
	for ptr = l; ptr.Next != nil; ptr = ptr.Next {
	}
	ptr.Next = NewLinkList(val)
}

func (l *LinkListNode[T]) InsertNext(val T) {
	if l == nil {
		return
	}

	newNode := NewLinkList(val)

	if l.Next != nil {
		nextNode := l.Next
		newNode.Next = nextNode
	}

	l.Next = newNode
}
