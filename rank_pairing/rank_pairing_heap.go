package rank_paring

import (
	"fmt"

	heap "github.com/theodesp/go-heaps"
)

type node struct {
	item               heap.Item
	left, next, parent *node
	rank               int
}

// RPHeap is an implementation of a rank Pairing Heap.
// The zero value for RPHeap Root is an empty Heap.
type RPHeap struct {
	head *node
	size int
}

type nInf struct{}

func (a nInf) Compare(b heap.Item) int {
	switch b.(type) {
	case nInf:
		return 0
	default:
		return -1
	}
}

// wrap compare function to support nInf (negative inf) type
func compare(a, b heap.Item) int {
	switch b.(type) {
	case nInf:
		switch a.(type) {
		case nInf:
			return 0
		default:
			return 1
		}
	default:
		return a.Compare(b)
	}
}

// Init initializes or clears the rankPairingHeap
func (p *RPHeap) Init() *RPHeap {
	p.head = &node{}
	p.size = 0
	return p
}

// New returns an initialized rankPairingHeap.
func New() *RPHeap { return new(RPHeap).Init() }

// FindMin returns the value of root
// Complexity: O(1)
func (r *RPHeap) FindMin() heap.Item {
	return r.head.item
}

// Insert the value val into the heap and return it
// Complexity: O(1)
func (r *RPHeap) Insert(val heap.Item) heap.Item {
	ptr := &node{
		item: val,
	}
	r.insertRoot(ptr)
	r.size++
	return val
}

// DeleteMin removes the top most value from the rankPairingHeap and returns it
// Complexity: O(log n)
func (r *RPHeap) DeleteMin() heap.Item {
	bucket := make([]*node, r.maxBucketSize())
	ret := r.head.item
	// an empty heap will panic here
	if ret == nil {
		return nil
	}
	r.size--
	for ptr := r.head.left; ptr != nil; {
		nextPtr := ptr.next
		ptr.next = nil
		ptr.parent = nil
		bucket = multiPass(bucket, ptr)
		ptr = nextPtr
	}
	for ptr := r.head.next; ptr != r.head; {
		nextPtr := ptr.next
		ptr.next = nil
		bucket = multiPass(bucket, ptr)
		ptr = nextPtr
	}
	r.head = &node{}
	for _, ptr := range bucket {
		if ptr != nil {
			r.insertRoot(ptr)
		}
	}
	return ret
}

// Clear the whole rankPairingHeap
func (r *RPHeap) Clear() {
	r.Init()
}

// Merge a rankPairingHeap r0 into a heap r, then clear r0
// Complexity: O(1)
func (r *RPHeap) Meld(a heap.Interface) heap.Interface {
	switch a.(type) {
	case *RPHeap:
	default:
		panic(fmt.Sprintf("unexpected type %T", a))
	}
	r0 := a.(*RPHeap)
	if r.head.item == nil {
		r.head = r0.head
		r.size = r0.size
		r0.Clear()
		return r
	}
	if r0.head.item == nil {
		return r
	}
	if compare(r.head.item, r0.head.item) < 0 {
		mergeRes := merge(r, r0)
		r0.Clear()
		return mergeRes
	} else {
		mergeRes := merge(r0, r)
		r0.Clear()
		return mergeRes
	}
}

// Size returns the size of the RPHeap
func (r *RPHeap) Size() int {
	return r.size
}

// Adjust the value of an item, since we have to find the item
// Complexity is O(n)
func (r *RPHeap) Adjust(old, new heap.Item) heap.Item {
	ptr := r.find(r.head, old)
	if ptr == nil {
		return nil
	}
	if compare(ptr.item, new) < 0 {
		r.decrease(ptr, nInf{})
		r.DeleteMin()
		r.Insert(new)
	} else {
		r.decrease(ptr, new)
	}
	return new
}

// Delete an item from the heap
// Complexity is O(n)
func (r *RPHeap) Delete(val heap.Item) heap.Item {
	ptr := r.find(r.head, val)
	if ptr == nil {
		return nil
	}
	if ptr == r.head {
		return r.DeleteMin()
	}
	r.decrease(ptr, nInf{})
	r.DeleteMin()
	return val
}

// Decrease the value of an item
// Complexity is O(log n)
func (r *RPHeap) decrease(ptr *node, val heap.Item) {
	if compare(val, ptr.item) < 0 {
		ptr.item = val
	}
	if ptr == r.head {
		return
	}
	if ptr.parent == nil {
		if compare(ptr.item, r.head.item) < 0 {
			r.head = ptr
		}
	} else {
		parent := ptr.parent
		if ptr == parent.left {
			parent.left = ptr.next
			if parent.left != nil {
				parent.left.parent = parent
			}
		} else {
			parent.next = ptr.next
			if parent.next != nil {
				parent.next.parent = parent
			}
		}
		ptr.next, ptr.parent = nil, nil
		if ptr.left != nil {
			ptr.rank = ptr.left.rank + 1
		} else {
			ptr.rank = 0
		}
		r.insertRoot(ptr)
		if parent.parent == nil {
			parent.rank = getrank(parent.left) + 1
		} else {
			for parent.parent != nil {
				leftrank := getrank(parent.left)
				nextrank := getrank(parent.next)
				newrank := leftrank + 1
				if leftrank != nextrank {
					if leftrank > nextrank {
						newrank = leftrank
					} else {
						newrank = nextrank
					}
				}
				if newrank >= parent.rank {
					break
				}
				parent.rank = newrank
				parent = parent.parent
			}
		}
	}
}

// Find the pointer to an item
// Complexity: O(n)
func (r *RPHeap) find(root *node, val heap.Item) *node {
	if root == nil {
		return nil
	} else if compare(root.item, val) == 0 {
		return root
	} else {
		if leftfind := r.find(root.left, val); leftfind != nil {
			return leftfind
		}
		for ptr := root.next; ptr != nil && ptr != root; ptr = ptr.next {
			if compare(ptr.item, val) == 0 {
				return ptr
			}
			if leftfind := r.find(ptr.left, val); leftfind != nil {
				return leftfind
			}
		}
	}
	return nil
}

func getrank(root *node) int {
	if root == nil {
		return -1
	}
	return root.rank
}

func merge(r0, r1 *RPHeap) *RPHeap {
	if r1.Size() == 1 {
		ptr := r1.head
		ptr.next, ptr.parent, ptr.left, ptr.rank = nil, nil, nil, 0
		r0.insertRoot(ptr)
		r0.size++
		return &RPHeap{
			head: r0.head,
			size: r0.size,
		}
	} else if r0.Size() == 1 {
		return merge(r1, r0)
	}
	r0.head.next, r1.head.next = r1.head.next, r0.head.next
	r0.size += r1.size
	return &RPHeap{
		head: r0.head,
		size: r0.size,
	}
}

func (r *RPHeap) maxBucketSize() int {
	bit, cnt := 1, r.size
	for cnt > 1 {
		cnt /= 2
		bit++
	}
	return bit + 1
}

func (r *RPHeap) insertRoot(ptr *node) {
	if r.head.item == nil {
		r.head = ptr
		ptr.next = ptr
	} else {
		ptr.next = r.head.next
		r.head.next = ptr
		if compare(ptr.item, r.head.item) < 0 {
			r.head = ptr
		}
	}
}

func multiPass(bucket []*node, ptr *node) []*node {
	for bucket[ptr.rank] != nil {
		rank := ptr.rank
		ptr = link(ptr, bucket[rank])
		bucket[rank] = nil
	}
	bucket[ptr.rank] = ptr
	return bucket
}

func link(left *node, right *node) *node {
	if right == nil {
		return left
	}
	var winner, loser *node
	if compare(right.item, left.item) < 0 {
		winner = right
		loser = left
	} else {
		winner = left
		loser = right
	}
	loser.parent = winner
	if winner.left != nil {
		loser.next = winner.left
		loser.next.parent = loser
	}
	winner.left = loser
	winner.rank = loser.rank + 1
	return winner
}
