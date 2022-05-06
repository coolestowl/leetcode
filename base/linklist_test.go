package base_test

import (
	"testing"

	"github.com/coolestowl/leetcode/base"
)

func TestLinkList(t *testing.T) {
	head := base.NewLinkList(-1)

	for i := range make([]struct{}, 10) {
		head.AppendTail(base.ElemType(i))
	}

	expected := base.ElemType(-1)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		if got := ptr.Val; got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
		expected++
	}

	for i := range make([]struct{}, 10) {
		head.InsertNext(base.ElemType(i))
	}

	innerSlice := []base.ElemType{-1, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	idx := 0
	for ptr := head; ptr != nil; ptr = ptr.Next {
		if got := ptr.Val; got != innerSlice[idx] {
			t.Errorf("expected %d, got %d", expected, got)
		}
		idx++
	}

	var nilHead *base.LinkListNode
	nilHead.AppendTail(233)
	nilHead.InsertNext(233)
	if nilHead != nil {
		t.Errorf("nil won't panic")
	}
}
