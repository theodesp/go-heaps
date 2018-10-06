// Package pairing implements a Pairing heap Data structure
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Pairing_heap
package pairing

import (
	heap "github.com/theodesp/go-heaps"
)

// PairHeap is an implementation of a Pairing Heap.
// The zero value for PairHeap Root is an empty Heap.
type PairHeap struct {
	root       *PairHeapNode
}

// PairHeapNode contains the current Value and the list if the sub-heaps
type PairHeapNode struct {
	// for use by client; untouched by this library
	Value heap.Item
	// List of children PairHeapNodes all containing values less than the Top of the heap
	children []*PairHeapNode
	// A reference to the parent Heap Node
	parent *PairHeapNode
	heap   *PairHeap // The heap to which this node belongs.
}

func (n *PairHeapNode) detach() []*PairHeapNode {
	if n.parent == nil {
		return nil // avoid detaching root
	}
	for _, node := range n.children {
		node.parent = n.parent
	}
	var idx int
	for i, node := range n.parent.children {
		if node == n {
			idx = i
			break
		}
	}
	n.parent.children = append(n.parent.children[:idx], n.parent.children[idx+1:]...)
	n.parent = nil
	return n.children
}

// Init initializes or clears the PairHeap
func (p *PairHeap) Init() *PairHeap {
	p.root = &PairHeapNode{}
	return p
}

// New returns an initialized PairHeap.
func New() *PairHeap { return new(PairHeap).Init() }

// IsEmpty returns true if PairHeap p is empty.
// The complexity is O(1).
func (p *PairHeap) IsEmpty() bool {
	return p.root.Value == nil
}

// Resets the current PairHeap
func (p *PairHeap) Clear() {
	p.root = &PairHeapNode{}
}

// Find the smallest item in the priority queue.
// The complexity is O(1).
func (p *PairHeap) FindMin() heap.Item {
	if p.IsEmpty() {
		return nil
	}
	return p.root.Value
}

// Inserts the value to the PairHeap and returns the Value
// The complexity is O(1).
func (p *PairHeap) Insert(v heap.Item) heap.Item {
	n := PairHeapNode{Value: v, heap: p}
	merge(&p.root, &n)
	return n.Value
}


// toDelete details what item to remove in a node call.
type toDelete int

const (
	removeItem toDelete = iota   // removes the given item
	removeMin                  // removes min item in the heap
)

// DeleteMin removes the top most value from the PairHeap and returns it
// The complexity is O(log n) amortized.
func (p *PairHeap) DeleteMin() heap.Item {
	return p.deleteItem(nil, removeMin)
}

// Deletes a PairHeapNode from the heap and returns the Value
// The complexity is O(log n) amortized.
func (p *PairHeap) Delete(item heap.Item) heap.Item {
	return p.deleteItem(item, removeItem)
}

func (p *PairHeap) deleteItem(item heap.Item, typ toDelete) heap.Item {
	var result PairHeapNode

	switch typ {
	case removeMin:
		if len(p.root.children) == 0 {
			result = *p.root
			p.root.Value = nil
		} else {
			result = *mergePairs(&p.root, p.root.children)
		}
	case removeItem:
		node := p.Find(item)
		if node == nil {
			return nil
		} else {
			children := node.detach()
			result = *mergePairs(&p.root, append(p.root.children, children...))
		}
	default:
		panic("invalid type")
	}
	return result.Value
}

// Adjusts the value to the PairHeapNode Value and returns it
// The complexity is O(n) amortized.
func (p *PairHeap) Adjust(item heap.Item, new heap.Item) heap.Item {
	node := p.Find(item)
	if node == nil {
		return nil
	}

	if node == p.root {
		p.DeleteMin()
		return p.Insert(new)
	} else {
		children := node.detach()
		node.Value = new
		mergePairs(&p.root, append(p.root.children, append([]*PairHeapNode{node}, children...)...))
		return node.Value
	}
}

// Do calls function cb on each element of the PairingHeap, in order of appearance.
// The behavior of Do is undefined if cb changes *p.
func (p *PairHeap) Do(cb func(item heap.Item)) {
	if p.IsEmpty() {
		return
	}
	// Call root first
	cb(p.root.Value)
	// Then continue to children
	visitChildren(p.root.children, cb)
}

// Exhausting search of the element that matches v. Returns it as a PairHeapNode
// The complexity is O(n) amortized.
func (p *PairHeap) Find(v heap.Item) *PairHeapNode {
	if p.IsEmpty() {
		return nil
	}

	if  p.root.Value.Compare(v) == 0 {
		return p.root
	} else {
		return p.findInChildren(p.root.children, v)
	}
}

func (p *PairHeap) findInChildren(children []*PairHeapNode, v heap.Item) *PairHeapNode {
	if len(children) == 0 {
		return nil
	}
	var node *PairHeapNode
loop:
	for _, heapNode := range children {
		cmp := heapNode.Value.Compare(v)
		switch {
		case cmp == 0: // found
			node = heapNode
			break loop
		default:
			node = p.findInChildren(heapNode.children, v)
			if node != nil {
				break loop
			}
		}
	}
	return node
}

func visitChildren(children []*PairHeapNode, cb func(item heap.Item)) {
	if len(children) == 0 {
		return
	}
	for _, heapNode := range children {
		cb(heapNode.Value)
		visitChildren(heapNode.children, cb)
	}
}

func merge(first **PairHeapNode, second *PairHeapNode) *PairHeapNode {
	q := *first
	if q.Value == nil { // Case when root is empty
		*first = second
		return *first
	}

	cmp := q.Value.Compare(second.Value)
	if cmp < 0 {
		// put 'second' as the first child of 'first' and update the parent
		q.children = append([]*PairHeapNode{second}, q.children...)
		second.parent = *first
		return *first
	} else {
		// put 'first' as the first child of 'second' and update the parent
		second.children = append([]*PairHeapNode{q}, second.children...)
		q.parent = second
		*first = second
		return second
	}
}

// Merges heaps together
func mergePairs(root **PairHeapNode, heaps []*PairHeapNode) *PairHeapNode {
	q := *root
	if len(heaps) == 1 {
		*root = heaps[0]
		heaps[0].parent = nil
		return q
	}
	var merged *PairHeapNode
	for { // iteratively merge heaps
		if len(heaps) == 0 {
			break
		}
		if merged == nil {
			merged = merge(&heaps[0], heaps[1])
			heaps = heaps[2:]
		} else {
			merged = merge(&merged, heaps[0])
			heaps = heaps[1:]
		}
	}
	*root = merged
	merged.parent = nil

	return q
}
