package sword2offer

import (
	"github.com/coolestowl/leetcode/base"
)

func LevelOrderI[T any](root *base.BinTreeNode[T]) []T {
	result := make([]T, 0)

	for queue, ptr, currentLen := []*base.BinTreeNode[T]{root}, 0, 1; ptr < len(queue); ptr, currentLen = currentLen, len(queue) {
		for idx := ptr; idx < currentLen; idx++ {
			node := queue[idx]
			if node == nil {
				continue
			}

			result, queue = append(result, node.Val), append(queue, node.Left, node.Right)
		}
	}

	return result
}
