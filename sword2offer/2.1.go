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
