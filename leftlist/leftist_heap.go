package leftlist

import (
	heap "github.com/theodesp/go-heaps"
)

// Node is a leaf in the heap.
type Node struct {
	Item        heap.Item
	Left, Right *Node
	s           int // s-value (or rank)
}

// LeftistHeap is a leftist heap implementation.
type LeftistHeap struct {
	Root *Node
}

func (h *LeftistHeap) merge(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	if x.Item.Compare(y.Item) == 1 {
		x, y = y, x
	}

	x.Right = h.merge(x.Right, y)

	if x.Left == nil {
		// left child doesn't exist, so move right child to the left side
		x.Left = x.Right
		x.Right = nil
	} else {
		// left child does exist, so compare s-values
		if x.Left.s < x.Right.s {
			x.Left, x.Right = x.Right, x.Left
		}
		// since we know the right child has the lower s-value, we can just
		// add one to its s-value
		x.s = x.Right.s + 1
	}

	return x
}

// Insert adds an item into the heap.
// The complexity is O(log n) amortized.
func (h *LeftistHeap) Insert(item heap.Item) heap.Item {
	h.Root = h.merge(&Node{
		Item: item,
	}, h.Root)

	return item
}

// DeleteMin deletes the minimum value and returns it.
// The complexity is O(log n) amortized.
func (h *LeftistHeap) DeleteMin() heap.Item {
	item := h.Root.Item

	h.Root = h.merge(h.Root.Left, h.Root.Right)

	return item
}

// FindMin finds the minimum value.
// The complexity is O(1).
func (h *LeftistHeap) FindMin() heap.Item {
	return h.Root.Item
}

// Clear removes all items from the heap.
func (h *LeftistHeap) Clear() {
	h.Root = nil
}
