// Package pairing implements a Pairing heap Data structure
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Pairing_heap
package skew

import (
	"sort"
	"testing"

	heap "github.com/theodesp/go-heaps"
)

func TestSkewHeapInteger(t *testing.T) {
	skew := &SkewHeap{}

	numbers := []int{4, 3, 2, 5}

	for _, number := range numbers {
		skew.Insert(Int(number))
	}

	sort.Ints(numbers)

	for _, number := range numbers {
		if Int(number) != skew.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
}

func TestSkewHeapString(t *testing.T) {
	skew := &SkewHeap{}

	strs := []string{"a", "ccc", "bb", "d"}

	for _, str := range strs {
		skew.Insert(Str(str))
	}

	sort.Strings(strs)

	for _, str := range strs {
		if Str(str) != skew.DeleteMin().(heap.String) {
			t.Fail()
		}
	}
}

func TestSkewHeap(t *testing.T) {
	skew := &SkewHeap{}

	numbers := []int{4, 3, -1, 5, 9}

	for _, number := range numbers {
		skew.Insert(Int(number))
	}

	if skew.FindMin() != Int(-1) {
		t.Fail()
	}

	skew.Clear()
	if skew.FindMin() != nil {
		t.Fail()
	}
}

func Int(value int) heap.Integer {
	return heap.Integer(value)
}

func Str(value string) heap.String {
	return heap.String(value)
}
