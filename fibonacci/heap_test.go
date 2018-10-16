package fibonacci

import (
	"sort"
	"testing"

	"github.com/theodesp/go-heaps"
)

func TestFibonacciHeapInteger(t *testing.T) {
	heap := &FibonacciHeap{}

	numbers := []int{4, 3, 2, 5}

	for _, number := range numbers {
		heap.Insert(Int(number))
	}

	sort.Ints(numbers)

	for _, number := range numbers {
		i := heap.DeleteMin().(go_heaps.Integer)
		if Int(number) != i {
			t.Fail()
		}
	}
}

func TestFibonacciHeapString(t *testing.T) {
	heap := &FibonacciHeap{}

	strs := []string{"a", "ccc", "bb", "d"}

	for _, str := range strs {
		heap.Insert(Str(str))
	}

	sort.Strings(strs)

	for _, str := range strs {
		if Str(str) != heap.DeleteMin().(go_heaps.String) {
			t.Fail()
		}
	}
}

func TestFibonacciHeap(t *testing.T) {
	heap := &FibonacciHeap{}

	numbers := []int{4, 3, -1, 5, 9}

	for _, number := range numbers {
		heap.Insert(Int(number))
	}

	if heap.FindMin() != Int(-1) {
		t.Fail()
	}

	heap.Clear()
	if heap.FindMin() != nil {
		t.Fail()
	}
}

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}

func Str(value string) go_heaps.String {
	return go_heaps.String(value)
}
