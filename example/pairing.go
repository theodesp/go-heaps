package main

import (
	"go-heaps/pairing"
	"fmt"
)

func main()  {
	heap := pairing.NewWithIntComparator()
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(5)

	fmt.Println(heap.DeleteMin()) // 2
	fmt.Println(heap.DeleteMin()) // 3
	fmt.Println(heap.DeleteMin()) // 4
	fmt.Println(heap.DeleteMin()) // 5
}
