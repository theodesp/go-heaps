package two_three

import "go-heaps"

// Package two_three implements a 2-3 heap Data structure
//
// Structure is not thread safe.
//
// Reference: http://www.cosc.canterbury.ac.nz/tad.takaoka/2-3heaps.pdf

// TwoThreeHeap represents a 2-3 Heap.
// The zero value for TwoThreeHeap Root is an empty Heap.
type TwoThreeHeap struct {
	Comparator go_heaps.Comparator // Key comparator
	children   []*TwoThreeHeapNode
}

// TwoThreeHeapNode contains the current Value and the references
type TwoThreeHeapNode struct {
	// for use by client; untouched by this library
	Value interface{}
	// References to relevant nodes
	parent, left, right, child  *TwoThreeHeapNode
}

// Init initializes or clears the TwoThreeHeap
func (p *TwoThreeHeap) Init(c go_heaps.Comparator) *TwoThreeHeap {
	p.children = []*TwoThreeHeapNode{}
	p.Comparator = c
	return p
}

// New returns an initialized TwoThreeHeap with the provided Comparator.
func New(c go_heaps.Comparator) *TwoThreeHeap { return new(TwoThreeHeap).Init(c) }

// NewWithIntComparator instantiates an TwoThreeHeap with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *TwoThreeHeap {
	return new(TwoThreeHeap).Init(go_heaps.IntComparator)
}

// NewWithStringComparator instantiates an TwoThreeHeap with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *TwoThreeHeap {
	return new(TwoThreeHeap).Init(go_heaps.StringComparator)
}

// IsEmpty returns true if TwoThreeHeap p is empty.
// The complexity is O(1).
func (p *TwoThreeHeap) IsEmpty() bool {
	return len(p.children) == 0
}

// Resets the current TwoThreeHeap
func (p *TwoThreeHeap) Clear() {
	p.children = []*TwoThreeHeapNode{}
}
