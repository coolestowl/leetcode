package base

type LinkListNode struct {
	Val  ElemType
	Next *LinkListNode
}

func NewLinkList(val ElemType) *LinkListNode {
	return &LinkListNode{
		Val:  val,
		Next: nil,
	}
}

func (l *LinkListNode) AppendTail(val ElemType) {
	if l == nil {
		return
	}

	var ptr *LinkListNode
	for ptr = l; ptr.Next != nil; ptr = ptr.Next {
	}
	ptr.Next = NewLinkList(val)
}

func (l *LinkListNode) InsertNext(val ElemType) {
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
