package leftistheap

import (
	"github.com/theodesp/go-heaps"
)

// Node is a leaf in the heap.
type Node struct {
	Item        go_heaps.Item
	Left, Right *Node

	s int
}

// LeftistHeap is a leftist heap implementation.
type LeftistHeap struct {
	Root *Node
}

func (heap *LeftistHeap) merge(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	if x.Item.Compare(y.Item) == 1 {
		x, y = y, x
	}

	x.Right = heap.merge(x.Right, y)

	if x.Left == nil {
		x.Left = x.Right
		x.Right = nil
	} else {
		if x.Left.s < x.Right.s {
			x.Left, x.Right = x.Right, x.Left
		}

		x.s = x.Right.s + 1
	}

	return x
}

// Insert adds an item into the heap.
func (heap *LeftistHeap) Insert(v go_heaps.Item) go_heaps.Item {
	heap.Root = heap.merge(&Node{
		Item: v,
	}, heap.Root)

	return v
}

// DeleteMin deletes the minimum value and returns it.
func (heap *LeftistHeap) DeleteMin() go_heaps.Item {
	item := heap.Root.Item

	heap.Root = heap.merge(heap.Root.Left, heap.Root.Right)

	return item
}

// FindMin finds the minimum value.
func (heap *LeftistHeap) FindMin() go_heaps.Item {
	return heap.Root.Item
}

// Clear removes all items from the heap.
func (heap *LeftistHeap) Clear() {
	heap.Root = nil
}
