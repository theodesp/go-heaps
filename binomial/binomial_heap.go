// Package binomial implements a Binomial heap Data structure
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Binomial_heap
package binomial

import (
	heap "github.com/theodesp/go-heaps"
)

// BinomialHeap is an implementation of a Binomial Heap.
type BinomialHeap struct {
	root *node
}

//node is a leaf in the heap
type node struct {
	item    heap.Item
	parent  *node
	child   *node
	sibling *node
	degree  int
}

//Insert inserts the value to the BinomialHeap and returns the item
// The complexity is O(log n).
func (b *BinomialHeap) Insert(v heap.Item) heap.Item {
	n := node{item: v}
	tempHeap := &BinomialHeap{root: &n}
	b.root = b.union(tempHeap)
	return n.item
}

// DeleteMin removes the smallest item from the BinomialHeap and returns it
// The complexity is O(log n) amortized.
func (b *BinomialHeap) DeleteMin() heap.Item {
	if (b.root == nil) {
		return nil
	}

	min := b.root
	var minPrev *node
	next := min.sibling
	nextPrev := min

	for (next != nil) {
		if (next.item.Compare(min.item) < 0) {
			min = next
			minPrev = nextPrev
		}
		nextPrev = next;
		next = next.sibling;
	}
	b.removeTreeRoot(min, minPrev)
	return min.item
}

// FindMin returns the smallest item in the heap.
// The complexity is O(1).
func (b *BinomialHeap) FindMin() heap.Item {
	if b.root == nil {
		return nil
	}
	min := b.root
	next := min.sibling

	for next != nil {
		if next.item.Compare(min.item) < 0 {
			min = next
		}
		next = next.sibling
	}

	return min.item
}

// Clear resets the current BinomialHeap
func (b *BinomialHeap) Clear() {
	b.root = nil
}

func (b *BinomialHeap) union(heap *BinomialHeap) *node {
	newRoot := merge(b, heap)
	b.root = nil
	heap.root = nil
	if newRoot == nil {
		return nil
	}
	var prev *node
	curr := newRoot
	next := newRoot.sibling
	for next != nil {
		if curr.degree != next.degree || (next.sibling != nil &&
			next.sibling.degree == curr.degree) {
			prev = curr
			curr = next
		} else {
			if curr.item.Compare(next.item) < 0 {
				curr.sibling = next.sibling
				linkNodes(curr, next)
			} else {
				if prev == nil {
					newRoot = next
				} else {
					prev.sibling = next
				}
				linkNodes(next, curr)
				curr = next
			}
		}
		next = curr.sibling
	}
	return newRoot
}

func merge(a *BinomialHeap, b *BinomialHeap) *node {
	if a.root == nil {
		return b.root
	}
	if b.root == nil {
		return a.root
	}

	var root *node
	aNext := a.root
	bNext := b.root
	if aNext.degree <= bNext.degree {
		root = aNext
		aNext = aNext.sibling
	} else {
		root = bNext
		bNext = bNext.sibling
	}

	tail := root

	for aNext != nil && bNext != nil {
		if aNext.degree <= bNext.degree {
			tail.sibling = aNext
			aNext = aNext.sibling
		} else {
			tail.sibling = bNext
			bNext = bNext.sibling
		}

		tail = tail.sibling
	}

	if aNext != nil {
		tail.sibling = aNext
	} else {
		tail.sibling = bNext
	}

	return root
}

func linkNodes(a, b *node) {
	b.parent = a
	b.sibling = a.child
	a.child = b
	a.degree++
}

func (b *BinomialHeap) removeTreeRoot(root, prev *node) {
	// Remove root from the heap
	if root == b.root {
		b.root = root.sibling
	} else {
		prev.sibling = root.sibling
	}

	var newRoot *node
	child := root.child
	for child != nil {
		next := child.sibling
		child.sibling = newRoot
		child.parent = nil
		newRoot = child
		child = next
	}
	newHeap := &BinomialHeap{root:newRoot}
	b.root = b.union(newHeap)
}
