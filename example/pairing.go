package main

import (
	"github.com/mb-14/go-heaps"
	"github.com/mb-14/go-heaps/pairing"
	"fmt"
)

func main()  {
	heap := pairing.New()
	heap.Insert(Int(4))
	heap.Insert(Int(3))
	heap.Insert(Int(2))
	heap.Insert(Int(5))

	fmt.Println(heap.DeleteMin()) // 2
	fmt.Println(heap.DeleteMin()) // 3
	fmt.Println(heap.DeleteMin()) // 4
	fmt.Println(heap.DeleteMin()) // 5
}

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}