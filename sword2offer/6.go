package sword2offer

import (
	"github.com/coolestowl/leetcode/base"
)

func LevelOrderI(root *base.BinTreeNode) []base.ElemType {
	result := make([]base.ElemType, 0)

	for queue, ptr, currentLen := []*base.BinTreeNode{root}, 0, 1; ptr < len(queue); ptr, currentLen = currentLen, len(queue) {
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

func LevelOrderII(root *base.BinTreeNode) [][]base.ElemType {
	result := make([][]base.ElemType, 0)

	for queue, ptr, currentLen := []*base.BinTreeNode{root}, 0, 1; ptr < len(queue); ptr, currentLen = currentLen, len(queue) {
		vals := make([]base.ElemType, 0, currentLen-ptr)

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

func LevelOrderIII(root *base.BinTreeNode) [][]base.ElemType {
	result, queue := make([][]base.ElemType, 0, 1), []*base.BinTreeNode{root}

	for reversed, ptr, currentLen := false, 0, 1; ptr < len(queue); reversed, ptr, currentLen = !reversed, currentLen, len(queue) {
		vals := make([]base.ElemType, 0, currentLen-ptr)

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
