package main

import (
	"fmt"

	"github.com/radlinskii/go-heaps/fibonacci"
	"github.com/theodesp/go-heaps"
)

func main() {
	heap1 := fibonacci.MakeHeap()
	heap1.Insert(&fibonacci.Node{Value: toInt(4)})
	heap1.Insert(&fibonacci.Node{Value: toInt(9)})
	heap1.Insert(&fibonacci.Node{Value: toInt(6)})

	heap2 := fibonacci.MakeHeap()
	heap2.Insert(&fibonacci.Node{Value: toInt(7)})
	heap2.Insert(&fibonacci.Node{Value: toInt(8)})
	heap2.Insert(&fibonacci.Node{Value: toInt(3)})
	heap2.Insert(&fibonacci.Node{Value: toInt(5)})

	fmt.Printf("heap1 - min: %v, n: %d\n", heap1.Minimum().Value, heap1.N)
	fmt.Printf("heap2 - min: %v, n: %d\n", heap2.Minimum().Value, heap2.N)

	heap3 := heap1.Union(heap2)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Value, heap3.N)
}

func toInt(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}
