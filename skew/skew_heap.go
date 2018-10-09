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
type Node struct {
	Item        heap.Item
	Right, Left *Node
}

func merge(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	if x.Item.Compare(y.Item) == 1 {
		x, y = y, x
	}

	x.Left, x.Right = x.Right, x.Left
	x.Left = merge(y, x.Left)

	return x
}

// SkewHeap is a skew heap implementation.
type SkewHeap struct {
	Root *Node
}

// Init initializes or clears the SkewHeap
func (h *SkewHeap) Init() *SkewHeap {
	return &SkewHeap{}
}

// New returns an initialized SkewHeap.
func New() *SkewHeap { return new(SkewHeap).Init() }

// Insert adds an item into the heap.
func (h *SkewHeap) Insert(v heap.Item) heap.Item {
	h.Root = merge(&Node{
		Item: v,
	}, h.Root)

	return v
}

// DeleteMin deletes the minimum value and returns it.
func (h *SkewHeap) DeleteMin() heap.Item {
	v := h.Root

	h.Root = merge(v.Right, v.Left)

	return v.Item
}

// FindMin finds the minimum value.
func (h *SkewHeap) FindMin() heap.Item {
	return h.Root.Item
}

// Clear removes all items from the heap.
func (h *SkewHeap) Clear() {
	h.Root = nil
}
