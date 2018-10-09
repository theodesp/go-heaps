package main

import (
	"fmt"

	"github.com/radlinskii/go-heaps/fibonacci"
	"github.com/theodesp/go-heaps"
)

func main() {
	heap1 := fibonacci.MakeHeap()
	heap1.Vis()
	heap1.Insert(&fibonacci.Node{Value: toInt(4)})
	heap1.Insert(&fibonacci.Node{Value: toInt(6)})
	heap1.Insert(&fibonacci.Node{Value: toInt(20)})
	heap1.Insert(&fibonacci.Node{Value: toInt(1)})

	heap2 := fibonacci.MakeHeap()
	heap2.Insert(&fibonacci.Node{Value: toInt(5)})
	heap2.Insert(&fibonacci.Node{Value: toInt(3)})
	heap2.Insert(&fibonacci.Node{Value: toInt(8)})
	heap2.Insert(&fibonacci.Node{Value: toInt(7)})
	heap2.Insert(&fibonacci.Node{Value: toInt(10)})

	fmt.Printf("heap1 - min: %v, n: %d\n", heap1.Minimum().Value, heap1.N)
	heap1.Vis()
	fmt.Printf("heap2 - min: %v, n: %d\n", heap2.Minimum().Value, heap2.N)
	heap2.Vis()

	heap3 := heap1.Union(heap2)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Value, heap3.N)
	heap3.Vis()

	n := heap3.ExtractMin()

	fmt.Printf("extractmin - v: %v\n", n.Value)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Value, heap3.N)

	heap3.Vis()

	n = heap3.ExtractMin()

	fmt.Printf("extractmin - v: %v\n", n.Value)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Value, heap3.N)

	heap3.Vis()
}

func toInt(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}
