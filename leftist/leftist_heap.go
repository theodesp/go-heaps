package leftist

import (
	heap "github.com/theodesp/go-heaps"
)

// Node is a leaf in the heap.
type Node struct {
	item        heap.Item
	left, right *Node
	s           int // s-value (or rank)
}

// LeftistHeap is a leftist heap implementation.
type LeftistHeap struct {
	root *Node
}

func mergeNodes(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}
	// Compare the roots of two heaps.
	if x.item.Compare(y.item) > 0 {
		return merge(y, x)
	} else {
		return merge(x, y)
	}
}

func merge(x, y *Node) *Node {
	if x.left == nil {
		// left child doesn't exist, so move right child to the smallest key
		// to maintain the leftList invariant
		x.left = y
		x.right = nil
	} else {
		x.right = mergeNodes(x.right, y)
		// left child does exist, so compare s-values
		if x.left.s < x.right.s {
			x.left, x.right = x.right, x.left
		}
		// since we know the right child has the lower s-value, we can just
		// add one to its s-value
		x.s = x.right.s + 1
	}

	return x
}

// Init initializes or clears the LeftistHeap
func (h *LeftistHeap) Init() *LeftistHeap {
	h.root = &Node{}
	return h
}

// New returns an initialized LeftistHeap.
func New() *LeftistHeap { return new(LeftistHeap).Init() }

// Insert adds an item into the heap.
// The complexity is O(log n) amortized.
func (h *LeftistHeap) Insert(item heap.Item) heap.Item {
	h.root = mergeNodes(&Node{
		item: item,
	}, h.root)

	return item
}

// DeleteMin deletes the minimum value and returns it.
// The complexity is O(log n) amortized.
func (h *LeftistHeap) DeleteMin() heap.Item {
	item := h.root.item

	h.root = mergeNodes(h.root.left, h.root.right)

	return item
}

// FindMin finds the minimum value.
// The complexity is O(1).
func (h *LeftistHeap) FindMin() heap.Item {
	return h.root.item
}

// Clear removes all items from the heap.
func (h *LeftistHeap) Clear() {
	h.Init()
}
