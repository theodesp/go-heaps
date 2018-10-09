package main

import (
	"fmt"

	"github.com/radlinskii/go-heaps/fibonacci" // after merging the pull request change it to "github.com/theodesp/go-heaps/fibonacci"
	goheap "github.com/theodesp/go-heaps"
)

func main() {
	heap1 := fibonacci.MakeHeap()
	heap1.Vis()
	heap1.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(4))})
	heap1.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(6))})
	heap1.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(20))})
	heap1.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(1))})

	heap2 := fibonacci.MakeHeap()
	node := &fibonacci.Node{Key: goheap.Item(goheap.Integer(5))}
	heap2.Insert(node)
	heap2.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(3))})
	node2 := &fibonacci.Node{Key: goheap.Item(goheap.Integer(8))}
	heap2.Insert(node2)
	heap2.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(7))})
	heap2.Insert(&fibonacci.Node{Key: goheap.Item(goheap.Integer(10))})

	fmt.Printf("heap1 - min: %v, n: %d\n", heap1.Minimum().Key, heap1.N)
	heap1.Vis()
	fmt.Printf("heap2 - min: %v, n: %d\n", heap2.Minimum().Key, heap2.N)
	heap2.Vis()

	heap3 := heap1.Union(heap2)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)
	heap3.Vis()

	n := heap3.ExtractMin()

	fmt.Printf("extractmin - v: %v\n", n.Key)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)

	heap3.Vis()

	n = heap3.ExtractMin()

	fmt.Printf("extractmin - v: %v\n", n.Key)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)

	heap3.Vis()

	heap3.DecreaseKey(node, goheap.Item(goheap.Integer(2)))
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)

	heap3.Vis()

	heap3.Delete(node2)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)

	heap3.Vis()

	heap4 := fibonacci.MakeHeap()
	heap4.Vis()
	heap4.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("igi"))})
	heap4.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("gigi"))})
	heap4.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("zigi"))})
	heap4.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("aga"))})

	heap5 := fibonacci.MakeHeap()
	node3 := &fibonacci.Node{Key: goheap.Item(goheap.String("iggi"))}
	heap5.Insert(node3)
	heap5.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("agga"))})
	node4 := &fibonacci.Node{Key: goheap.Item(goheap.String("agzi"))}
	heap5.Insert(node4)
	heap5.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("iaga"))})
	heap5.Insert(&fibonacci.Node{Key: goheap.Item(goheap.String("iga"))})

	fmt.Printf("heap4 - min: %v, n: %d\n", heap4.Minimum().Key, heap4.N)
	heap4.Vis()
	fmt.Printf("heap5 - min: %v, n: %d\n", heap5.Minimum().Key, heap5.N)
	heap5.Vis()

	heap6 := heap4.Union(heap5)
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)
	heap6.Vis()

	n2 := heap6.ExtractMin()

	fmt.Printf("extractmin - v: %v\n", n2.Key)
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)

	heap6.Vis()

	n2 = heap6.ExtractMin()

	fmt.Printf("extractmin - v: %v\n", n2.Key)
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)

	heap6.Vis()

	heap6.DecreaseKey(node3, goheap.Item(goheap.String("agu")))
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)

	heap6.Vis()

	heap6.Delete(node4)
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)

	heap6.Vis()
}
