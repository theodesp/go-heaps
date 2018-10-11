package main

import (
	"fmt"

	goheap "github.com/theodesp/go-heaps"
	"github.com/theodesp/go-heaps/treap"
)

type Int int

func (a Int) Compare(b goheap.Item) int {
	a1 := a
	a2 := b.(Int)
	switch {
	case a1 > a2:
		return 1
	case a1 < a2:
		return -1
	default:
		return 0
	}
}

func (i Int) Print() {
	fmt.Println(i)
}

func main() {
	heap := treap.New()
	heap.Insert(Int(4))
	heap.Insert(Int(19))
	heap.Insert(Int(8))
	heap.Insert(Int(27))
	heap.Insert(Int(20))
	heap.Insert(Int(12))
	heap.Insert(Int(15))
	heap.Insert(Int(6))
	heap.Insert(Int(7))
	heap.Insert(Int(8))

	fmt.Println(heap.DeleteMin()) // 4
	fmt.Println(heap.DeleteMin()) // 6
	fmt.Println(heap.DeleteMin()) // 7
	fmt.Println(heap.DeleteMin()) // 8
}
