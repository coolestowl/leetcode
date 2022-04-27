package base

import (
	"fmt"

	"github.com/m1gwings/treedrawer/tree"
)

type BinTreeNode[T any] struct {
	Val   T
	Left  *BinTreeNode[T]
	Right *BinTreeNode[T]
}

/*
[1, 2, 3, 4, -1, 5, 6, -1, 7]

to

      1
	2    3
  4     5  6
   7
*/
func NewBinTreeFromSlice[T any](slice []T, isEmpty func(T) bool) *BinTreeNode[T] {
	return nonRecursive(slice, isEmpty)
	// return recursive(slice, 0, isEmpty)
}

func recursive[T any](slice []T, idx int, isEmpty func(T) bool) *BinTreeNode[T] {
	if idx >= len(slice) || isEmpty(slice[idx]) {
		return nil
	}

	return &BinTreeNode[T]{
		Val:   slice[idx],
		Left:  recursive(slice, idx*2+1, isEmpty),
		Right: recursive(slice, idx*2+2, isEmpty),
	}
}

func nonRecursive[T any](slice []T, isEmpty func(T) bool) *BinTreeNode[T] {
	if len(slice) == 0 {
		return nil
	}

	nodes := make([]*BinTreeNode[T], 0, len(slice))
	for idx := range slice {
		if isEmpty(slice[idx]) {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &BinTreeNode[T]{
				Val: slice[idx],
			})
		}
	}

	for idx := range nodes {
		if nodes[idx] == nil {
			continue
		}

		if leftIdx := idx*2 + 1; leftIdx < len(nodes) && nodes[leftIdx] != nil {
			nodes[idx].Left = nodes[leftIdx]
		}
		if rightIdx := idx*2 + 2; rightIdx < len(nodes) && nodes[rightIdx] != nil {
			nodes[idx].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}

func DrawBinTree[T fmt.Stringer](root *BinTreeNode[T]) *tree.Tree {
	tr := tree.NewTree(tree.NodeString(root.Val.String()))
	drawBinTreeRecursively(tr, root)
	return tr
}

func drawBinTreeRecursively[T fmt.Stringer](parent *tree.Tree, node *BinTreeNode[T]) {
	if left := node.Left; left != nil {
		tr := parent.AddChild(tree.NodeString(left.Val.String()))
		drawBinTreeRecursively(tr, left)
	}
	if right := node.Right; right != nil {
		tr := parent.AddChild(tree.NodeString(right.Val.String()))
		drawBinTreeRecursively(tr, right)
	}
}

func PreOrder[A, B any](root *BinTreeNode[A], f func(*BinTreeNode[A]) B) []B {
	result := make([]B, 0)
	preorder(root, f, &result)
	return result
}

func preorder[A, B any](root *BinTreeNode[A], f func(*BinTreeNode[A]) B, output *[]B) {
	if root == nil {
		return
	}

	*output = append(*output, f(root))
	preorder(root.Left, f, output)
	preorder(root.Right, f, output)
}

func InOrder[A, B any](root *BinTreeNode[A], f func(*BinTreeNode[A]) B) []B {
	result := make([]B, 0)
	inorder(root, f, &result)
	return result
}

func inorder[A, B any](root *BinTreeNode[A], f func(*BinTreeNode[A]) B, output *[]B) {
	if root == nil {
		return
	}

	inorder(root.Left, f, output)
	*output = append(*output, f(root))
	inorder(root.Right, f, output)
}

func PostOrder[A, B any](root *BinTreeNode[A], f func(*BinTreeNode[A]) B) []B {
	result := make([]B, 0)
	postorder(root, f, &result)
	return result
}

func postorder[A, B any](root *BinTreeNode[A], f func(*BinTreeNode[A]) B, output *[]B) {
	if root == nil {
		return
	}

	postorder(root.Left, f, output)
	postorder(root.Right, f, output)
	*output = append(*output, f(root))
}
