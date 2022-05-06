package sword2offer

import "github.com/coolestowl/leetcode/base"

func ReversePrint(head *base.LinkListNode) []base.ElemType {
	s := base.NewStack()
	for ptr := head; ptr != nil; ptr = ptr.Next {
		s.Push(ptr.Val)
	}

	result := make([]base.ElemType, 0, s.Len())
	for s.Len() > 0 {
		result = append(result, s.Pop())
	}
	return result
}

func Reverse(head *base.LinkListNode) *base.LinkListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, cur := (*base.LinkListNode)(nil), head

	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}

	return pre
}

type Node struct {
	Val    base.ElemType
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	return NodeItemsToLinkList(LinkListToNodeItems(head))
}

type NodeItem struct {
	Val       base.ElemType
	RandomIdx int
}

func LinkListToNodeItems(head *Node) []NodeItem {
	count, ptrIdxMap := 0, make(map[*Node]int)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		ptrIdxMap[ptr] = count
		count++
	}

	result := make([]NodeItem, 0, count)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		randomIdx := -1
		if ptr.Random != nil {
			randomIdx = ptrIdxMap[ptr.Random]
		}

		result = append(result, NodeItem{
			Val:       ptr.Val,
			RandomIdx: randomIdx,
		})
	}

	return result
}

func NodeItemsToLinkList(items []NodeItem) *Node {
	if len(items) == 0 {
		return nil
	}

	var (
		h     *Node = nil
		tail        = h
		nodes       = make([]*Node, len(items))
	)

	for i, item := range items {
		newNode := &Node{
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
