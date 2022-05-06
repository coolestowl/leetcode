package base_test

import (
	"strconv"
	"testing"

	"github.com/coolestowl/leetcode/base"
)

type stringerInt int

func (i stringerInt) String() string {
	return strconv.Itoa(int(i))
}

func TestBinaryTree(t *testing.T) {
	root := base.NewBinTreeFromSlice([]base.ElemType{1, 2, 3, 3, 4, 5, 6, -1, -1, 2, 3, 3, 3}, func(t base.ElemType) bool { return t == -1 })

	tr := base.DrawBinTree(root)

	t.Log(tr)
}
