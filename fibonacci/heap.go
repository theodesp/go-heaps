// Package fibonacci implements a Fibonacci Heap Data structure
// Reference: https://en.wikipedia.org/wiki/Fibonacci_heap
package fibonacci

import (
	"fmt"

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
		addNode(fh.Min, x)

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
	fh2.Min.left, newFH.Min.left = newFH.Min.left, fh2.Min.left

	if fh.Min == nil || (fh2.Min != nil && fh.Min.Value.Compare(fh2.Min.Value) > 0) {
		newFH.Min = fh2.Min
	}
	newFH.N = fh.N + fh2.N
	return newFH
}

// ExtractMin extracts the node with minimum value from a heap
// and returns pointer to this node.
func (fh *Heap) ExtractMin() *Node {
	z := fh.Min
	if z != nil {
		for {
			if c := z.child; c != nil {
				c.parent = nil
				if c.right != c {
					z.child = c.right
					c.right.left = c.left
					c.left.right = c.right
				} else {
					z.child = nil
				}
				c.left = z.left
				c.right = z
				z.left.right = c
				z.left = c
			} else {
				break
			}
		}

		z.left.right = z.right
		z.right.left = z.left

		if z == z.right {
			fh.Min = nil
		} else {
			fh.Min = z.right
			fh.consolidate()
		}
		fh.N--
	}
	return z
}

func (fh *Heap) consolidate() {
	degreeToRoot := make(map[int]*Node)
	w := fh.Min
	last := w.left
	for {
		r := w.right                                         // wchodze dla 3 dla 3 nic nie zrobie wiec wstawie 3 do mapy
		x := w                                               // wchodze dla 5 dla 5 wejde w if bo jest cos w mapie
		d := x.degree                                        // w mapie jest 3 z 5 jako dziecko na 1 a na 0 nic nie ma
		fmt.Println(" XXXX ", x.Value, " xr", x.right.Value) // wchdoze dla 9 podstawiam 9 pod 0 w mapie
		for {                                                // wchodze dla 2 2 najpierw spotka 9 a pozniej 3 w mapie i jest zapisana na d:2 w mapie z dziecmi 3 i 9
			if y, ok := degreeToRoot[d]; !ok { // wchodze dla 4 i wrzucam ja do mapy na 0
				fmt.Println(" x ", x.Value, d, degreeToRoot) // wchodze dla 6 i patrze ze jest 4 na 0 wiec robie 6 dzieckiem 4
				break                                        // w mapie jest 4 na 1 i 2 na 2
			} else { // wchodze dla 8 i wstawiam ja na 0 w mapie
				fmt.Println(" YYYY ", y.Value)    // TERAZ POWINIENEM SKONCZYC, 3 powinna byc kolejna po 8 ale jest 2 !!!
				if y.Value.Compare(x.Value) < 0 { // 2 nie jest trojka wiec petla sie nie konczy tylko idzie dalej
					y, x = x, y // czemu 8 ma wskaznik na 2 po prawej stronie ?
				}
				fh.link(y, x)
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
	fmt.Println("MAAAP", degreeToRoot)
	fh.Min = nil
	for _, v := range degreeToRoot {
		fh.addRoot(v)
	}

}

func (fh *Heap) link(y, x *Node) {
	y.right.left = y.left
	y.left.right = y.right

	y.parent = x

	if x.child == nil {
		x.child = y
		y.left = y
		y.right = y
	} else {
		y.left = x.child.left
		y.right = x.child
		x.child.left.right = y
		x.child.left = y
	}

	y.mark = false
}

func (fh *Heap) addRoot(n *Node) {
	if fh.Min == nil {
		n.left = n
		n.right = n
		fh.Min = n
	} else {
		addNode(fh.Min, n)
		if n.Value.Compare(fh.Min.Value) < 0 {
			fh.Min = n
		}
	}
}

func addNode(h, x *Node) {
	h.left.right = x
	x.right = h
	x.left = h.left
	h.left = x
}

// Vis visualize
func (fh Heap) Vis() {
	if fh.Min == nil {
		fmt.Println("<empty>")
		return
	}
	var f func(*Node, string)
	f = func(n *Node, pre string) {
		pc := "│ "
		for x := n; ; x = x.right {
			if x.right != n {
				fmt.Print(pre, "├─")
			} else {
				fmt.Print(pre, "└─")
				pc = "  "
			}
			if x.child == nil {
				fmt.Println("╴", x.Value)
			} else {
				fmt.Println("┐", x.Value)
				f(x.child, pre+pc)
			}
			if x.right == n {
				break
			}
		}
	}
	f(fh.Min, "")
}

func (n *Node) GetRight() *Node {
	return n.right
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) GetChild() *Node {
	return n.child
}
