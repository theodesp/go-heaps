// Package fibonacci implements a Fibonacci Heap Data structure
// Reference: https://en.wikipedia.org/wiki/Fibonacci_heap
package fibonacci

import (
	heap "github.com/theodesp/go-heaps"
)

// Heap is a implementation of Fibonacci heap.
type Heap struct {
	Min *Node
	N   int
}

// Node holds structure of nodes inside Fibonacci heap.
type Node struct {
	Value                      heap.Item
	left, right, parent, child *Node
	mark                       bool
	degree                     int
}

// MakeHeap creates and returns a new, empty heap.
func MakeHeap() *Heap {
	var fh Heap
	fh.Min = nil
	fh.N = 0
	return &fh
}

// Insert inserts a new node, with predeclared value, to the heap.
func (fh *Heap) Insert(x *Node) *Node {
	x.degree = 0
	x.mark = false
	x.parent = nil
	x.child = nil

	if fh.Min == nil {
		x.left = x
		x.right = x
		fh.Min = x
	} else {
		fh.Min.left.right = x
		x.right = fh.Min
		x.left = fh.Min.left
		fh.Min.left = x

		if fh.Min.Value.Compare(x.Value) > 0 {
			fh.Min = x
		}
	}
	fh.N++
	return x
}

// Minimum returns pointer to the heap's node holding the minimum value.
func (fh *Heap) Minimum() *Node {
	return fh.Min
}

// Union creates and returns the merge of two mergeable heaps.
func (fh *Heap) Union(fh2 *Heap) *Heap {
	newFH := MakeHeap()
	newFH.Min = fh.Min

	newFH.Min.left.right = fh2.Min
	fh2.Min.left.right = newFH.Min
	fh2.Min.left = newFH.Min.left
	newFH.Min.left = fh2.Min.left

	if fh.Min == nil || (fh2.Min != nil && fh.Min.Value.Compare(fh2.Min.Value) > 0) {
		newFH.Min = fh2.Min
	}
	newFH.N = fh.N + fh2.N
	return newFH
}
