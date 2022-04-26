package base_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/coolestowl/leetcode/base"
	"github.com/m1gwings/treedrawer/tree"
)

type stringerInt int

func (i stringerInt) String() string {
	return strconv.Itoa(int(i))
}

func TestBinaryTree(t *testing.T) {
	root := base.NewBinTreeFromSlice([]stringerInt{1, 2, 3, 3, 4, 5, 6, -1, -1, 2, 3, 3, 3}, func(t stringerInt) bool { return t == -1 })

	tr := drawBinTree(root)

	t.Log(tr)
}

func drawBinTree[T fmt.Stringer](root *base.BinTreeNode[T]) *tree.Tree {
	tr := tree.NewTree(tree.NodeString(root.Val.String()))
	drawBinTreeRecursively(tr, root)
	return tr
}

func drawBinTreeRecursively[T fmt.Stringer](parent *tree.Tree, node *base.BinTreeNode[T]) {
	if left := node.Left; left != nil {
		tr := parent.AddChild(tree.NodeString(left.Val.String()))
		drawBinTreeRecursively(tr, left)
	}
	if right := node.Right; right != nil {
		tr := parent.AddChild(tree.NodeString(right.Val.String()))
		drawBinTreeRecursively(tr, right)
	}
}
