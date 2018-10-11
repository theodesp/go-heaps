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

func (this Node) CompareKey (than *Node) int {
    return this.Key.Compare(than.Key)
}

func (this Node) ComparePriority (than *Node) int {
    return this.Priority.Compare(than.Priority)
}

func merge(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
    }
    
    if x.ComparePriority(y) != 0 {
        if x.ComparePriority(y) == 1 {
            x, y = y, x
        }

        if x.CompareKey(y) == 1 {
            x.Left = merge(x.Left, y)
            return x
        } else {
            x.Right = merge(x.Right, y)
            return x
        }
    }

	if x.CompareKey(y) == 1 {
		x, y = y, x
	}

	x.Left, x.Right = x.Right, x.Left
	x.Left = merge(y, x.Left)

	return x
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
func (h *Treap) Init() *Treap {
	return &Treap{}
}

// New returns an initialized Treap.
func New() *Treap { return new(Treap).Init() }

// Insert adds an item into the heap.
func (h *Treap) Insert(v goheap.Item) goheap.Item {
	h.Root = merge(&Node{
        Priority: generatePriority(),
        Key: v,
	}, h.Root)

	return v
}

// DeleteMin deletes the minimum value and returns it.
func (h *Treap) DeleteMin() goheap.Item {
    v := h.Root
    if v.Left == nil {
        h.Root = v.Right
        return v.Key
    }

    for ; v.Left.Left != nil; v = v.Left {}
    
    min := v.Left
    v.Left = nil
    return min.Key
}

func (n *Node) findMinNode() *Node {
    if n == nil {
        return nil
    }

    minL := n.Left.findMinNode()
    minR := n.Right.findMinNode()

    if minL == nil {
		return minR
	}

	if minR == nil {
		return minL
    }

    if minL.CompareKey(minR) == 1 {
        minL, minR = minR, minL
    }

    if minL.CompareKey(n) == 1 {
        return n
    } else {
        return minL
    }
}

// FindMin finds the minimum value.
func (h *Treap) FindMin() goheap.Item {
    v := h.Root
    for ; v.Left != nil; v = v.Left {}
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
    fmt.Printf("%d-",h.Priority)
    h.Key.(Printable).Print()
    h.Right.Print(lv + 1)
}