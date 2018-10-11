// Package fibonacci implements a Fibonacci FibonacciHeap Data structure
// Reference: https://en.wikipedia.org/wiki/Fibonacci_heap
// Implementation from Introduction to Algorithms by T. Cormen
package fibonacci

import (
	goheap "github.com/theodesp/go-heaps"
)

// FibonacciHeap is a implementation of Fibonacci heap.
type FibonacciHeap struct {
	root   *Node
	length int
}

// Node holds structure of nodes inside Fibonacci heap.
type Node struct {
	item                      goheap.Item
	prev, next, parent, child *Node
	isMarked                  bool
	degree                    int
}

func (fh *FibonacciHeap) insertRoot(x *Node) {
	if fh.root == nil {
		// create fh's root list containing only x
		x.prev = x
		x.next = x
		fh.root = x
	} else {
		// insert x to fh's root list
		insert(fh.root, x)
		if x.item.Compare(fh.root.item) < 0 {
			fh.root = x
		}
	}
}

func insert(x, y *Node) {
	x.prev.next = y
	y.next = x
	y.prev = x.prev
	x.prev = y
}

// New creates and returns a new, empty heap.
func New() *FibonacciHeap {
	return &FibonacciHeap{root: nil, length: 0}
}

// Insert inserts a new node, with predeclared item, to the heap.
func (fh *FibonacciHeap) Insert(v goheap.Item) goheap.Item {
	x := &Node{item: v, isMarked: false}

	fh.insertRoot(x)
	fh.length++
	return v
}

// FindMin returns pointer to the heap's node holding the minimum item.
func (fh *FibonacciHeap) FindMin() goheap.Item {
	return fh.root.item
}

// DeleteMin extracts the node with minimum item from a heap
// and returns pointer to this node.
func (fh *FibonacciHeap) DeleteMin() goheap.Item {
	z := fh.root
	if z == nil {
		return nil
	}
	for {
		// add z children to fh's root list
		if x := z.child; x != nil {
			x.parent = nil
			if x.next != x {
				z.child = x.next
				x.next.prev = x.prev
				x.prev.next = x.next
			} else {
				z.child = nil
			}
			x.prev = z.prev
			x.next = z
			z.prev.next = x
			z.prev = x
		} else {
			break
		}
	}
	// remove z from fh's root list
	z.prev.next = z.next
	z.next.prev = z.prev

	if z == z.next {
		fh.root = nil
	} else {
		fh.root = z.next
		fh.consolidate()
	}
	fh.length--

	return z.item
}

func (fh *FibonacciHeap) consolidate() {
	degreeToRoot := make(map[int]*Node)
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

func link(x, y *Node) {
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

// Clear resets heap.
func (fh *FibonacciHeap) Clear() {
	fh.length = 0
	fh.root = &Node{}
}

/*
// DecreaseKey decreases the key of given node.
func (fh *FibonacciHeap) DecreaseKey(x *Node, k goheap.Item) {
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

func (fh *FibonacciHeap) cut(x, y *Node) {
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

func (fh *FibonacciHeap) cascadingCut(y *Node) {
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

/*
// Delete deletes node x from heap fh.
func (fh *FibonacciHeap) Delete(x *Node) {
	switch x.item.(type) {
	case goheap.Integer:
		fh.DecreaseKey(x, goheap.Item(goheap.Integer(-1<<63)))
	case goheap.String:
		fh.DecreaseKey(x, goheap.Item(goheap.String("")))
	}
	fh.DeleteMin()
}
*/
