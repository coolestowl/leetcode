package base

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
