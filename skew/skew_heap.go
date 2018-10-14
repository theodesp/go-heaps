// Package skew implements a Skew heap Data structure
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Skew_heap
package skew

import (
	heap "github.com/theodesp/go-heaps"
)

// Node is a leaf in the heap.
type node struct {
	item        heap.Item
	right, left *node
}

func merge(x, y *node) *node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}
	// x should point to the smaller item
	if x.item.Compare(y.item) > 0 {
		x, y = y, x
	}

	x.left, x.right = x.right, x.left
	x.left = merge(y, x.left)

	return x
}

// SkewHeap is a skew heap implementation.
type SkewHeap struct {
	root *node
}

// Init initializes or clears the SkewHeap
func (h *SkewHeap) Init() *SkewHeap {
	return &SkewHeap{}
}

// New returns an initialized SkewHeap.
func New() *SkewHeap { return new(SkewHeap).Init() }

// Insert adds an item into the heap.
func (h *SkewHeap) Insert(v heap.Item) heap.Item {
	h.root = merge(&node{
		item: v,
	}, h.root)

	return v
}

// DeleteMin deletes the minimum value and returns it.
func (h *SkewHeap) DeleteMin() heap.Item {
	v := h.root

	h.root = merge(v.right, v.left)

	return v.item
}

// FindMin finds the minimum value.
func (h *SkewHeap) FindMin() heap.Item {
	return h.root.item
}

// Clear removes all items from the heap.
func (h *SkewHeap) Clear() {
	h.Init()
}
