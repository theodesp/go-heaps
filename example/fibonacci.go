package main

import (
	"fmt"

	"github.com/radlinskii/go-heaps/fibonacci"
)

func main() {
	heap1 := fibonacci.MakeHeap()
	heap1.Vis()
	heap1.Insert(&fibonacci.Node{Key: fibonacci.ToInt(4)})
	heap1.Insert(&fibonacci.Node{Key: fibonacci.ToInt(6)})
	heap1.Insert(&fibonacci.Node{Key: fibonacci.ToInt(20)})
	heap1.Insert(&fibonacci.Node{Key: fibonacci.ToInt(1)})

	heap2 := fibonacci.MakeHeap()
	node := &fibonacci.Node{Key: fibonacci.ToInt(5)}
	heap2.Insert(node)
	heap2.Insert(&fibonacci.Node{Key: fibonacci.ToInt(3)})
	node2 := &fibonacci.Node{Key: fibonacci.ToInt(8)}
	heap2.Insert(node2)
	heap2.Insert(&fibonacci.Node{Key: fibonacci.ToInt(7)})
	heap2.Insert(&fibonacci.Node{Key: fibonacci.ToInt(10)})

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

	heap3.DecreaseKey(node, fibonacci.ToInt(2))
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)

	heap3.Vis()

	heap3.Delete(node2)
	fmt.Printf("heap3 - min: %v, n: %d\n", heap3.Minimum().Key, heap3.N)

	heap3.Vis()

	heap4 := fibonacci.MakeHeap()
	heap4.Vis()
	heap4.Insert(&fibonacci.Node{Key: fibonacci.ToString("igi")})
	heap4.Insert(&fibonacci.Node{Key: fibonacci.ToString("gigi")})
	heap4.Insert(&fibonacci.Node{Key: fibonacci.ToString("zigi")})
	heap4.Insert(&fibonacci.Node{Key: fibonacci.ToString("aga")})

	heap5 := fibonacci.MakeHeap()
	node3 := &fibonacci.Node{Key: fibonacci.ToString("iggi")}
	heap5.Insert(node3)
	heap5.Insert(&fibonacci.Node{Key: fibonacci.ToString("agga")})
	node4 := &fibonacci.Node{Key: fibonacci.ToString("agzi")}
	heap5.Insert(node4)
	heap5.Insert(&fibonacci.Node{Key: fibonacci.ToString("iaga")})
	heap5.Insert(&fibonacci.Node{Key: fibonacci.ToString("iga")})

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

	heap6.DecreaseKey(node3, fibonacci.ToString("agu"))
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)

	heap6.Vis()

	heap6.Delete(node4)
	fmt.Printf("heap6 - min: %v, n: %d\n", heap6.Minimum().Key, heap6.N)

	heap6.Vis()
}
