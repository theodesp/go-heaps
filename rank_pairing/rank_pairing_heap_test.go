package rank_paring

import (
	"sort"
	"testing"

	heap "github.com/theodesp/go-heaps"
)

func TestRPHeapInteger(t *testing.T) {
	rpheap := &RPHeap{}
	numbers := []int{4, 3, 2, 5}
	for _, number := range numbers {
		rpheap.Insert(Int(number))
	}
	sort.Ints(numbers)
	for _, number := range numbers {
		if Int(number) != rpheap.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
}

func TestRPHeapAdjustDecrease(t *testing.T) {
	rpheap := &RPHeap{}
	numbers := []int{4, 3, 2, 5}
	for _, number := range numbers {
		rpheap.Insert(Int(number))
	}
	ans := []int{1, 2, 3, 5}
	rpheap.Adjust(Int(4), Int(1))
	for _, number := range ans {
		if Int(number) != rpheap.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
}

func TestRPHeapAdjustIncrease(t *testing.T) {
	rpheap := &RPHeap{}
	numbers := []int{4, 3, 2, 5}
	for _, number := range numbers {
		rpheap.Insert(Int(number))
	}
	ans := []int{2, 3, 5, 6}
	rpheap.Adjust(Int(4), Int(6))
	for _, number := range ans {
		if Int(number) != rpheap.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
}

func TestRPHeapDelete(t *testing.T) {
	rpheap := &RPHeap{}
	numbers := []int{4, 3, 2, 5}
	for _, number := range numbers {
		rpheap.Insert(Int(number))
	}
	ans := []int{2, 3, 5}
	rpheap.Delete(Int(4))
	for _, number := range ans {
		if Int(number) != rpheap.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
}

func TestRPHeapString(t *testing.T) {
	rpheap := &RPHeap{}

	strs := []string{"a", "ccc", "bb", "d"}

	for _, str := range strs {
		rpheap.Insert(Str(str))
	}

	sort.Strings(strs)

	for _, str := range strs {
		if Str(str) != rpheap.DeleteMin().(heap.String) {
			t.Fail()
		}
	}
}

func TestRPHeapInteger2(t *testing.T) {
	rpheap := &RPHeap{}
	rpheap.Insert(Int(5))
	rpheap.Insert(Int(3))
	if rpheap.DeleteMin().(heap.Integer) != Int(3) {
		t.Fail()
	}
	rpheap.Insert(Int(4))
	rpheap.Insert(Int(2))
	if rpheap.DeleteMin().(heap.Integer) != Int(2) {
		t.Fail()
	}
	if rpheap.DeleteMin().(heap.Integer) != Int(4) {
		t.Fail()
	}
	if rpheap.DeleteMin().(heap.Integer) != Int(5) {
		t.Fail()
	}
}

func TestRPHeapMerge0(t *testing.T) {
	runTestMerge([]int{2, 8, 5, 7}, []int{4, 9, 6}, t)
}

func TestRPHeapMerge1(t *testing.T) {
	runTestMerge([]int{4, 9, 6}, []int{2, 8, 5, 7}, t)
}

func TestRPHeapMerge2(t *testing.T) {
	runTestMerge([]int{2}, []int{4, 9, 6}, t)
}

func TestRPHeapMerge3(t *testing.T) {
	runTestMerge([]int{2, 8, 5, 7}, []int{4}, t)
}

func runTestMerge(arr1, arr2 []int, t *testing.T) {
	ans := append(arr1, arr2...)
	sort.Ints(ans)
	rpheap1 := &RPHeap{}
	rpheap2 := &RPHeap{}
	for _, number := range arr1 {
		rpheap1.Insert(Int(number))
	}
	for _, number := range arr2 {
		rpheap2.Insert(Int(number))
	}
	rpheap1 = rpheap1.Meld(rpheap2).(*RPHeap)
	for _, number := range ans {
		if Int(number) != rpheap1.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
	if rpheap2.Size() != 0 {
		t.Error("rpheap2 is not empty")
		t.Fail()
	}
	if rpheap1.Size() != 0 {
		t.Error("rpheap1 is not empty: size =", rpheap1.Size())
		t.Fail()
	}
}

func Int(value int) heap.Integer {
	return heap.Integer(value)
}

func Str(value string) heap.String {
	return heap.String(value)
}
