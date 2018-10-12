// Package skew implements a Treap (tree-heap) Data structure
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Treap

package treap

import (
	"fmt"
	"math/rand"

	goheap "github.com/theodesp/go-heaps"
)

type Node struct {
	Priority    goheap.Integer
	Key         goheap.Item
	Left, Right *Node
}

// Split treap into 2 treaps:
// - All key in left treap <= key
// - All key in right treap > key
func split(t *Node, key goheap.Item) (*Node, *Node) {
	var (
		left, right *Node
	)

	if t == nil {
		return nil, nil
	} else if t.Key.Compare(key) != 1 {
		t.Right, right = split(t.Right, key)
		left := t
		return left, right
	} else {
		left, t.Left = split(t.Left, key)
		right := t
		return left, right
	}
}

// Merge 2 treaps into one with condition:
// max key on left treap is <= than min key on right treap
func merge(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	if x.Priority.Compare(y.Priority) == 1 {
		x.Right = merge(x.Right, y)
		return x
	} else {
		y.Left = merge(x, y.Left)
		return y
	}
}

func (t *Node) insert(pnode *Node) *Node {
	if t == nil {
		return pnode
	}

	if pnode.Priority.Compare(t.Priority) == 1 {
		pnode.Left, pnode.Right = split(t, pnode.Key)
		return pnode
	}

	if t.Key.Compare(pnode.Key) != 1 {
		t.Right = t.Right.insert(pnode)
	} else {
		t.Left = t.Left.insert(pnode)
	}
	return t
}

// Generate priority for new node.
func generatePriority() goheap.Integer {
	return goheap.Integer(rand.Intn(65536))
}

// Treap implementation.
type Treap struct {
	Root *Node
}

// Init initializes or clears the Treap
func (h *Treap) Init() *Treap { return &Treap{} }

// New returns an initialized Treap.
func New() *Treap { return new(Treap).Init() }

// Insert adds an item into the heap.
func (h *Treap) Insert(v goheap.Item) goheap.Item {
	pnode := &Node{
		Priority: generatePriority(),
		Key:      v,
	}

	if h.Root == nil {
		h.Root = pnode
	} else {
		h.Root = h.Root.insert(pnode)
	}
	return v
}

// DeleteMin deletes the minimum value and returns it.
func (h *Treap) DeleteMin() goheap.Item {
	v := h.Root
	if v.Left == nil {
		h.Root = v.Right
		return v.Key
	}

	for ; v.Left.Left != nil; v = v.Left {
	}

	min := v.Left
	v.Left = nil
	return min.Key
}

// FindMin finds the minimum value.
func (h *Treap) FindMin() goheap.Item {
	v := h.Root
	for ; v.Left != nil; v = v.Left {
	}
	return v.Key
}

// Clear removes all items from the heap.
func (h *Treap) Clear() {
	h.Root = nil
}

type Printable interface {
	Print()
}

func (h *Node) Print(lv int) {
	if h == nil {
		return
	}

	h.Left.Print(lv + 1)
	for i := 0; i < lv; i++ {
		fmt.Print("-")
	}
	fmt.Printf("%d-", h.Priority)
	h.Key.(Printable).Print()
	h.Right.Print(lv + 1)
}
