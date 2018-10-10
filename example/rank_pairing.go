package main

import (
	"fmt"

	"github.com/theodesp/go-heaps"
	rpheap "github.com/theodesp/go-heaps/rank_pairing"
)

func main() {
	heap := rpheap.New()
	heap.Insert(Int(4))
	heap.Insert(Int(3))
	heap.Insert(Int(2))
	heap.Insert(Int(5))

	fmt.Println(heap.DeleteMin()) // 2
	fmt.Println(heap.DeleteMin()) // 3
	fmt.Println(heap.DeleteMin()) // 4
	fmt.Println(heap.DeleteMin()) // 5

	heap1 := rpheap.New()
	heap2 := rpheap.New()
	heap1.Insert(Int(2))
	heap1.Insert(Int(8))
	heap1.Insert(Int(5))
	heap1.Insert(Int(7))
	heap2.Insert(Int(4))
	heap2.Insert(Int(9))
	heap2.Insert(Int(6))

	heap1.Merge(heap2)

	fmt.Println(heap1.DeleteMin()) // 2
	fmt.Println(heap1.DeleteMin()) // 4
	fmt.Println(heap1.DeleteMin()) // 5
	//...
}

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}
