package main

import (
	"fmt"

	goheap "github.com/theodesp/go-heaps"
	"github.com/theodesp/go-heaps/treap"
)

func main() {
	heap := treap.New()
	heap.Insert(goheap.Integer(4))
	heap.Insert(goheap.Integer(19))
	heap.Insert(goheap.Integer(8))
	heap.Insert(goheap.Integer(27))
	heap.Insert(goheap.Integer(20))
	heap.Insert(goheap.Integer(12))
	heap.Insert(goheap.Integer(15))
	heap.Insert(goheap.Integer(6))
	heap.Insert(goheap.Integer(7))
	heap.Insert(goheap.Integer(8))

	fmt.Println(heap.DeleteMin()) // 4
	fmt.Println(heap.DeleteMin()) // 6
	fmt.Println(heap.DeleteMin()) // 7
	fmt.Println(heap.DeleteMin()) // 8
	fmt.Println(heap.DeleteMin()) // 8
	fmt.Println(heap.DeleteMin()) // 12
	fmt.Println(heap.DeleteMin()) // 15
	fmt.Println(heap.DeleteMin()) // 19
	fmt.Println(heap.DeleteMin()) // 20
	fmt.Println(heap.DeleteMin()) // 27
}
