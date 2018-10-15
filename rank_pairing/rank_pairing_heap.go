package rank_paring

import (
	"fmt"

	heap "github.com/theodesp/go-heaps"
)

type node struct {
	Val                heap.Item
	Left, Next, Parent *node
	Rank               int
}

// RPHeap is an implementation of a Rank Pairing Heap.
// The zero value for RPHeap Root is an empty Heap.
type RPHeap struct {
	head *node
	size int
}

type NInf struct{}

func (a NInf) Compare(b heap.Item) int {
	switch b.(type) {
	case NInf:
		return 0
	default:
		return -1
	}
}

// wrap compare function to support NInf (negative inf) type
func compare(a, b heap.Item) int {
	switch b.(type) {
	case NInf:
		switch a.(type) {
		case NInf:
			return 0
		default:
			return 1
		}
	default:
		return a.Compare(b)
	}
}

// Init initializes or clears the RankPairingHeap
func (p *RPHeap) Init() *RPHeap {
	return &RPHeap{}
}

// New returns an initialized RankPairingHeap.
func New() *RPHeap { return new(RPHeap).Init() }

// FindMin returns the value of root
// Complexity: O(1)
func (r *RPHeap) FindMin() heap.Item {
	return r.head.Val
}

// Insert the value val into the heap and return it
// Complexity: O(1)
func (r *RPHeap) Insert(val heap.Item) heap.Item {
	ptr := &node{
		Val: val,
	}
	r.insertRoot(ptr)
	r.size++
	return val
}

// DeleteMin removes the top most value from the RankPairingHeap and returns it
// Complexity: O(log n)
func (r *RPHeap) DeleteMin() heap.Item {
	bucket := make([]*node, r.maxBucketSize())
	ret := r.head.Val
	// an empty heap will panic here
	r.size--
	for ptr := r.head.Left; ptr != nil; {
		nextPtr := ptr.Next
		ptr.Next = nil
		ptr.Parent = nil
		bucket = multiPass(bucket, ptr)
		ptr = nextPtr
	}
	for ptr := r.head.Next; ptr != r.head; {
		nextPtr := ptr.Next
		ptr.Next = nil
		bucket = multiPass(bucket, ptr)
		ptr = nextPtr
	}
	r.head = nil
	for _, ptr := range bucket {
		if ptr != nil {
			r.insertRoot(ptr)
		}
	}
	return ret
}

// Clear the whole RankPairingHeap
func (r *RPHeap) Clear() {
	r.head = nil
	r.size = 0
}

// Merge a RankPairingHeap r0 into a heap r, then clear r0
// Complexity: O(1)
func (r *RPHeap) Meld(a heap.Interface) heap.Interface {
	switch a.(type) {
	case *RPHeap:
	default:
		panic(fmt.Sprintf("unexpected type %T", a))
	}
	r0 := a.(*RPHeap)
	if r.head == nil {
		r.head = r0.head
		r.size = r0.size
		r0.Clear()
		return r
	}
	if r0.head == nil {
		return r
	}
	if compare(r.head.Val, r0.head.Val) < 0 {
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
	if compare(ptr.Val, new) < 0 {
		r.decrease(ptr, NInf{})
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
	r.decrease(ptr, NInf{})
	r.DeleteMin()
	return val
}

// Decrease the value of an item
// Complexity is O(log n)
func (r *RPHeap) decrease(ptr *node, val heap.Item) {
	if compare(val, ptr.Val) < 0 {
		ptr.Val = val
	}
	if ptr == r.head {
		return
	}
	if ptr.Parent == nil {
		if compare(ptr.Val, r.head.Val) < 0 {
			r.head = ptr
		}
	} else {
		parent := ptr.Parent
		if ptr == parent.Left {
			parent.Left = ptr.Next
			if parent.Left != nil {
				parent.Left.Parent = parent
			}
		} else {
			parent.Next = ptr.Next
			if parent.Next != nil {
				parent.Next.Parent = parent
			}
		}
		ptr.Next, ptr.Parent = nil, nil
		if ptr.Left != nil {
			ptr.Rank = ptr.Left.Rank + 1
		} else {
			ptr.Rank = 0
		}
		r.insertRoot(ptr)
		if parent.Parent == nil {
			parent.Rank = getRank(parent.Left) + 1
		} else {
			for parent.Parent != nil {
				leftRank := getRank(parent.Left)
				nextRank := getRank(parent.Next)
				newRank := leftRank + 1
				if leftRank != nextRank {
					if leftRank > nextRank {
						newRank = leftRank
					} else {
						newRank = nextRank
					}
				}
				if newRank >= parent.Rank {
					break
				}
				parent.Rank = newRank
				parent = parent.Parent
			}
		}
	}
}

// Find the pointer to an item
// Complexity: O(n)
func (r *RPHeap) find(root *node, val heap.Item) *node {
	if root == nil {
		return nil
	} else if compare(root.Val, val) == 0 {
		return root
	} else {
		if leftfind := r.find(root.Left, val); leftfind != nil {
			return leftfind
		}
		for ptr := root.Next; ptr != nil && ptr != root; ptr = ptr.Next {
			if compare(ptr.Val, val) == 0 {
				return ptr
			}
			if leftfind := r.find(ptr.Left, val); leftfind != nil {
				return leftfind
			}
		}
	}
	return nil
}

func getRank(root *node) int {
	if root == nil {
		return -1
	}
	return root.Rank
}

func merge(r0, r1 *RPHeap) *RPHeap {
	if r1.Size() == 1 {
		ptr := r1.head
		ptr.Next, ptr.Parent, ptr.Left, ptr.Rank = nil, nil, nil, 0
		r0.insertRoot(ptr)
		r0.size++
		return &RPHeap{
			head: r0.head,
			size: r0.size,
		}
	} else if r0.Size() == 1 {
		return merge(r1, r0)
	}
	r0.head.Next, r1.head.Next = r1.head.Next, r0.head.Next
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
	if r.head == nil {
		r.head = ptr
		ptr.Next = ptr
	} else {
		ptr.Next = r.head.Next
		r.head.Next = ptr
		if compare(ptr.Val, r.head.Val) < 0 {
			r.head = ptr
		}
	}
}

func multiPass(bucket []*node, ptr *node) []*node {
	for bucket[ptr.Rank] != nil {
		rank := ptr.Rank
		ptr = link(ptr, bucket[rank])
		bucket[rank] = nil
	}
	bucket[ptr.Rank] = ptr
	return bucket
}

func link(left *node, right *node) *node {
	if right == nil {
		return left
	}
	var winner, loser *node
	if compare(right.Val, left.Val) < 0 {
		winner = right
		loser = left
	} else {
		winner = left
		loser = right
	}
	loser.Parent = winner
	if winner.Left != nil {
		loser.Next = winner.Left
		loser.Next.Parent = loser
	}
	winner.Left = loser
	winner.Rank = loser.Rank + 1
	return winner
}
