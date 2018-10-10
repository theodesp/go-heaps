package rank_paring

import (
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
func (r *RPHeap) Merge(r0 *RPHeap) {
	if r.head == nil {
		r.head = r0.head
		r.size = r0.size
		r0.Clear()
		return
	}
	if r0.head == nil {
		return
	}
	if r.head.Val.Compare(r0.head.Val) < 0 {
		mergeRes := merge(r, r0)
		r.size = mergeRes.size
		r.head = mergeRes.head
	} else {
		mergeRes := merge(r0, r)
		r.size = mergeRes.size
		r.head = mergeRes.head
	}
	r0.Clear()
}

// Size returns the size of the RPHeap
func (r *RPHeap) Size() int {
	return r.size
}

func merge(r0, r1 *RPHeap) *RPHeap {
	if r1.Size() == 1 {
		ptr := r1.head
		ptr.Next, ptr.Parent, ptr.Left, ptr.Rank = nil, nil, nil, 0
		r0.insertRoot(ptr)
		r0.size++
		return r0
	} else if r0.Size() == 1 {
		return merge(r1, r0)
	}
	r1.head.Next.Parent = r0.head
	r0.head.Next.Parent = r1.head.Next
	r0.head.Next, r1.head.Next = r1.head.Next, r0.head.Next
	r0.size += r1.size
	return r0
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
		if ptr.Val.Compare(r.head.Val) < 0 {
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
	if right.Val.Compare(left.Val) < 0 {
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
