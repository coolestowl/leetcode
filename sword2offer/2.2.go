package sword2offer

import "github.com/coolestowl/leetcode/base"

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
