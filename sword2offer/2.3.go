package sword2offer

type Node[T any] struct {
	Val    T
	Next   *Node[T]
	Random *Node[T]
}

func CopyRandomList[T any](head *Node[T]) *Node[T] {
	return NodeItemsToLinkList(LinkListToNodeItems(head))
}

type NodeItem[T any] struct {
	Val       T
	RandomIdx int
}

func LinkListToNodeItems[T any](head *Node[T]) []NodeItem[T] {
	count, ptrIdxMap := 0, make(map[*Node[T]]int)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		ptrIdxMap[ptr] = count
		count++
	}

	result := make([]NodeItem[T], 0, count)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		randomIdx := -1
		if ptr.Random != nil {
			randomIdx = ptrIdxMap[ptr.Random]
		}

		result = append(result, NodeItem[T]{
			Val:       ptr.Val,
			RandomIdx: randomIdx,
		})
	}

	return result
}

func NodeItemsToLinkList[T any](items []NodeItem[T]) *Node[T] {
	if len(items) == 0 {
		return nil
	}

	var (
		h     *Node[T] = nil
		tail           = h
		nodes          = make([]*Node[T], len(items))
	)

	for i, item := range items {
		newNode := &Node[T]{
			Val: item.Val,
		}

		nodes[i] = newNode

		if tail != nil {
			tail.Next = newNode
		} else {
			h = newNode
		}
		tail = newNode
	}

	idx := 0
	for ptr := h; ptr != nil; ptr, idx = ptr.Next, idx+1 {
		item := items[idx]
		if item.RandomIdx > -1 {
			ptr.Random = nodes[item.RandomIdx]
		}
	}

	return h
}
