package sword2offer

import "github.com/coolestowl/leetcode/base"

func LevelOrderIII[T any](root *base.BinTreeNode[T]) [][]T {
	result, queue := make([][]T, 0, 1), []*base.BinTreeNode[T]{root}

	for reversed, ptr, currentLen := false, 0, 1; ptr < len(queue); reversed, ptr, currentLen = !reversed, currentLen, len(queue) {
		vals := make([]T, 0, currentLen-ptr)

		for idx := currentLen - 1; idx >= ptr; idx-- {
			node := queue[idx]
			if node == nil {
				continue
			}

			vals = append(vals, node.Val)
			if reversed {
				queue = append(queue, node.Right, node.Left)
			} else {
				queue = append(queue, node.Left, node.Right)
			}
		}

		if len(vals) > 0 {
			result = append(result, vals)
		}
	}

	return result
}
