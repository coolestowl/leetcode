package sword2offer

import "github.com/coolestowl/leetcode/base"

func ReversePrint[T Elem](head *base.LinkListNode[T]) []T {
	s := base.NewStack[T]()
	for ptr := head; ptr != nil; ptr = ptr.Next {
		s.Push(ptr.Val)
	}

	result := make([]T, 0, s.Len())
	for s.Len() > 0 {
		result = append(result, s.Pop())
	}
	return result
}

func Reverse[T Elem](head *base.LinkListNode[T]) *base.LinkListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}

	pre, cur := (*base.LinkListNode[T])(nil), head

	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}

	return pre
}

type Node[T any] struct {
	Val    T
	Next   *Node[T]
	Random *Node[T]
}

func CopyRandomList[T any](head *Node[T]) *Node[T] {
	return NodeItemsToLinkList(LinkListToNodeItems(head))
}

type NodeItem[T any] struct {
	Val       T
	RandomIdx int
}

func LinkListToNodeItems[T any](head *Node[T]) []NodeItem[T] {
	count, ptrIdxMap := 0, make(map[*Node[T]]int)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		ptrIdxMap[ptr] = count
		count++
	}

	result := make([]NodeItem[T], 0, count)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		randomIdx := -1
		if ptr.Random != nil {
			randomIdx = ptrIdxMap[ptr.Random]
		}

		result = append(result, NodeItem[T]{
			Val:       ptr.Val,
			RandomIdx: randomIdx,
		})
	}

	return result
}

func NodeItemsToLinkList[T any](items []NodeItem[T]) *Node[T] {
	if len(items) == 0 {
		return nil
	}

	var (
		h     *Node[T] = nil
		tail           = h
		nodes          = make([]*Node[T], len(items))
	)

	for i, item := range items {
		newNode := &Node[T]{
			Val: item.Val,
		}

		nodes[i] = newNode

		if tail != nil {
			tail.Next = newNode
		} else {
			h = newNode
		}
		tail = newNode
	}

	idx := 0
	for ptr := h; ptr != nil; ptr, idx = ptr.Next, idx+1 {
		item := items[idx]
		if item.RandomIdx > -1 {
			ptr.Random = nodes[item.RandomIdx]
		}
	}

	return h
}
