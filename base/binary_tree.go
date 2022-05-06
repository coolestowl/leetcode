package base

import (
	"github.com/m1gwings/treedrawer/tree"
)

type BinTreeNode struct {
	Val   ElemType
	Left  *BinTreeNode
	Right *BinTreeNode
}

/*
[1, 2, 3, 4, -1, 5, 6, -1, 7]

to

      1
	2    3
  4     5  6
   7
*/
func NewBinTreeFromSlice(slice []ElemType, isEmpty func(ElemType) bool) *BinTreeNode {
	return nonRecursive(slice, isEmpty)
	// return recursive(slice, 0, isEmpty)
}

func recursive(slice []ElemType, idx int, isEmpty func(ElemType) bool) *BinTreeNode {
	if idx >= len(slice) || isEmpty(slice[idx]) {
		return nil
	}

	return &BinTreeNode{
		Val:   slice[idx],
		Left:  recursive(slice, idx*2+1, isEmpty),
		Right: recursive(slice, idx*2+2, isEmpty),
	}
}

func nonRecursive(slice []ElemType, isEmpty func(ElemType) bool) *BinTreeNode {
	if len(slice) == 0 {
		return nil
	}

	nodes := make([]*BinTreeNode, 0, len(slice))
	for idx := range slice {
		if isEmpty(slice[idx]) {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &BinTreeNode{
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

func DrawBinTree(root *BinTreeNode) *tree.Tree {
	tr := tree.NewTree(tree.NodeString(root.Val.String()))
	drawBinTreeRecursively(tr, root)
	return tr
}

func drawBinTreeRecursively(parent *tree.Tree, node *BinTreeNode) {
	if left := node.Left; left != nil {
		tr := parent.AddChild(tree.NodeString(left.Val.String()))
		drawBinTreeRecursively(tr, left)
	}
	if right := node.Right; right != nil {
		tr := parent.AddChild(tree.NodeString(right.Val.String()))
		drawBinTreeRecursively(tr, right)
	}
}

func PreOrder(root *BinTreeNode, f func(*BinTreeNode)) {
	if root == nil {
		return
	}

	f(root)
	PreOrder(root.Left, f)
	PreOrder(root.Right, f)
}

func InOrder(root *BinTreeNode, f func(*BinTreeNode)) {
	if root == nil {
		return
	}

	InOrder(root.Left, f)
	f(root)
	InOrder(root.Right, f)
}

func PostOrder(root *BinTreeNode, f func(*BinTreeNode)) {
	if root == nil {
		return
	}

	PostOrder(root.Left, f)
	PostOrder(root.Right, f)
	f(root)
}
