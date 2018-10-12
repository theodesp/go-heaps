// Package fibonacci implements a Fibonacci FibonacciHeap Data structure
// Reference: https://en.wikipedia.org/wiki/Fibonacci_heap
// Implementation from Introduction to Algorithms by T. Cormen
package fibonacci

import (
	heap "github.com/theodesp/go-heaps"
)

// FibonacciHeap is a implementation of Fibonacci heap.
type FibonacciHeap struct {
	root   *node
}

// node holds structure of nodes inside Fibonacci heap.
type node struct {
	item                      heap.Item
	prev, next, parent, child *node
	isMarked                  bool
	degree                    int
}

// New creates and returns a new, empty heap.
func New() *FibonacciHeap {
	return &FibonacciHeap{root: nil}
}

// Insert inserts a new node, with predeclared item, to the heap.
func (fh *FibonacciHeap) Insert(item heap.Item) heap.Item {
	n := &node{item: item, isMarked: false}

	fh.insertRoot(n)
	return item
}

// FindMin returns the minimum item.
func (fh *FibonacciHeap) FindMin() heap.Item {
	if fh.root == nil {
		return nil
	}
	return fh.root.item
}

// DeleteMin extracts the node with minimum item from a heap
// and returns the minimum item.
func (fh *FibonacciHeap) DeleteMin() heap.Item {
	r := fh.root
	if r == nil {
		return nil
	}
	for {
		// add r children to fh's root list
		if x := r.child; x != nil {
			x.parent = nil
			if x.next != x {
				r.child = x.next
				x.next.prev = x.prev
				x.prev.next = x.next
			} else {
				r.child = nil
			}
			x.prev = r.prev
			x.next = r
			r.prev.next = x
			r.prev = x
		} else {
			break
		}
	}
	// remove r from fh's root list
	r.prev.next = r.next
	r.next.prev = r.prev

	if r == r.next {
		fh.root = nil
	} else {
		fh.root = r.next
		fh.consolidate()
	}

	return r.item
}

func (fh *FibonacciHeap) consolidate() {
	degreeToRoot := make(map[int]*node)
	w := fh.root
	last := w.prev
	for {
		r := w.next
		x := w
		d := x.degree
		for {
			if y, ok := degreeToRoot[d]; !ok {
				break
			} else {
				if y.item.Compare(x.item) < 0 {
					y, x = x, y
				}
				link(x, y)
				delete(degreeToRoot, d)
				d++
			}
		}
		degreeToRoot[d] = x
		if w == last {
			break
		}
		w = r
	}
	fh.root = nil
	for _, v := range degreeToRoot {
		fh.insertRoot(v)
	}

}

// Clear resets heap.
func (fh *FibonacciHeap) Clear() {
	fh.root = nil
}

func link(x, y *node) {
	// remove y from fh's root list
	y.next.prev = y.prev
	y.prev.next = y.next
	// make y a child of x and increase degree of x
	y.parent = x
	if x.child == nil {
		x.child = y
		y.prev = y
		y.next = y
	} else {
		insert(x.child, y)
	}

	y.isMarked = false
}

func (fh *FibonacciHeap) insertRoot(n *node) {
	if fh.root == nil {
		// create fh's root list containing only n
		n.prev = n
		n.next = n
		fh.root = n
	} else {
		// insert n to fh's root list
		insert(fh.root, n)
		if n.item.Compare(fh.root.item) < 0 {
			fh.root = n
		}
	}
}

func insert(x, y *node) {
	x.prev.next = y
	y.next = x
	y.prev = x.prev
	x.prev = y
}

/*
// DecreaseKey decreases the key of given node.
func (fh *FibonacciHeap) DecreaseKey(x *node, k heap.Item) {
	if x.item.Compare(k) < 0 {
		panic("new item is greater than the previous one")
	}
	x.item = k
	y := x.parent
	if y != nil && x.item.Compare(y.item) < 0 {
		fh.cut(x, y)
		fh.cascadingCut(y)
	}
	if x.item.Compare(fh.root.item) < 0 {
		fh.root = x
	}
}

func (fh *FibonacciHeap) cut(x, y *node) {
	// remove x from y's children list and decrement y's degree
	if x.next != x {
		y.child = x.next
		x.next.prev = x.prev
		x.prev.next = x.next
	} else {
		y.child = nil
	}
	y.degree--
	// add x to fh's root list
	insert(fh.root, x)

	x.parent = nil
	x.isMarked = false
}

func (fh *FibonacciHeap) cascadingCut(y *node) {
	z := y.parent
	if z != nil {
		if !y.isMarked {
			y.isMarked = true
		} else {
			fh.cut(y, z)
			fh.cascadingCut(z)
		}
	}
}
*/
