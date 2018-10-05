package two_three

// Package two_three implements a 2-3 heap Data structure
//
// Structure is not thread safe.
//
// Reference: http://www.cosc.canterbury.ac.nz/tad.takaoka/2-3heaps.pdf

// TwoThreeHeap represents a 2-3 Heap.
// The zero value for TwoThreeHeap Root is an empty Heap.
type TwoThreeHeap struct {
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
func (p *TwoThreeHeap) Init() *TwoThreeHeap {
	p.children = []*TwoThreeHeapNode{}
	return p
}

// New returns an initialized TwoThreeHeap
func New() *TwoThreeHeap { return new(TwoThreeHeap).Init() }

// IsEmpty returns true if TwoThreeHeap p is empty.
// The complexity is O(1).
func (p *TwoThreeHeap) IsEmpty() bool {
	return len(p.children) == 0
}

// Resets the current TwoThreeHeap
func (p *TwoThreeHeap) Clear() {
	p.children = []*TwoThreeHeapNode{}
}
