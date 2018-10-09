package main

import (
	"github.com/theodesp/go-heaps"
	"github.com/theodesp/go-heaps/skew"
	"fmt"
)

func main() {
	heap := skewheap.New()
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

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}
