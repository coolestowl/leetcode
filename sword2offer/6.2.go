package sword2offer

import (
	"github.com/coolestowl/leetcode/base"
)

func LevelOrderII[T any](root *base.BinTreeNode[T]) [][]T {
	result := make([][]T, 0)

	for queue, ptr, currentLen := []*base.BinTreeNode[T]{root}, 0, 1; ptr < len(queue); ptr, currentLen = currentLen, len(queue) {
		vals := make([]T, 0, currentLen-ptr)

		for idx := ptr; idx < currentLen; idx++ {
			node := queue[idx]
			if node == nil {
				continue
			}

			vals, queue = append(vals, node.Val), append(queue, node.Left, node.Right)
		}

		if len(vals) > 0 {
			result = append(result, vals)
		}
	}

	return result
}
