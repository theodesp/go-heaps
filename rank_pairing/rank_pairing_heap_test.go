package rank_paring

import (
	"sort"
	"testing"

	heap "github.com/theodesp/go-heaps"
)

func TestRPHeapInteger(t *testing.T) {
	rpheap := New()
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
	if rpheap.DeleteMin() != nil {
		t.Fail()
	}
	if rpheap.FindMin() != nil {
		t.Fail()
	}
}

func TestRPHeapAdjustDecrease(t *testing.T) {
	rpheap := New()
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

func TestRPHeapAdjustDecrease2(t *testing.T) {
	rpheap := New()
	numbers := []int{9, 2, 4, 3, 1, 5, 6, 8, 7, 0}
	for _, number := range numbers {
		rpheap.Insert(Int(number))
	}
	rpheap.DeleteMin()
	ans := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	rpheap.Adjust(Int(4), Int(0))
	for _, number := range ans {
		if res := rpheap.DeleteMin().(heap.Integer); Int(number) != res {
			t.Fail()
		}
	}
}

func TestRPHeapAdjustNotExist(t *testing.T) {
	rpheap := New()
	numbers := []int{4, 3, 2, 5}
	for _, number := range numbers {
		rpheap.Insert(Int(number))
	}
	ans := []int{2, 3, 4, 5}
	rpheap.Adjust(Int(6), Int(1))
	for _, number := range ans {
		if Int(number) != rpheap.DeleteMin().(heap.Integer) {
			t.Fail()
		}
	}
}

func TestRPHeapAdjustIncrease(t *testing.T) {
	rpheap := New()
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
	rpheap := New()
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
	rpheap := New()

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
	rpheap := New()
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

func TestRPHeapMerge(t *testing.T) {
	runTestMerge([]int{2, 8, 5, 7}, []int{4, 9, 6}, t)
	runTestMerge([]int{4, 9, 6}, []int{2, 8, 5, 7}, t)
	runTestMerge([]int{2}, []int{4, 9, 6}, t)
	runTestMerge([]int{2, 8, 5, 7}, []int{4}, t)
	runTestMerge([]int{2, 8, 5, 7}, []int{}, t)
	runTestMerge([]int{}, []int{4, 9, 6}, t)
}

func runTestMerge(arr1, arr2 []int, t *testing.T) {
	ans := append(arr1, arr2...)
	sort.Ints(ans)
	rpheap1 := New()
	rpheap2 := New()
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
