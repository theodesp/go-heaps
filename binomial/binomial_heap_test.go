package binomial

import (
	"sort"
	"testing"

	"github.com/theodesp/go-heaps"
)

func TestLeftistHeapInteger(t *testing.T) {
	heap := &BinomialHeap{}

	numbers := []int{4, 3, 2, 5}

	for _, number := range numbers {
		heap.Insert(Int(number))
	}

	sort.Ints(numbers)

	for _, number := range numbers {
		if Int(number) != heap.DeleteMin().(go_heaps.Integer) {
			t.Fail()
		}
	}
}

func TestBinomialHeapIntegerDelete(t *testing.T) {
	heap := &BinomialHeap{}

	numbers := []int{14, 112, 15, 16 , 71, 91, 1, 12, 23, 56, 34}
	
	for _, number := range numbers {
		heap.Insert(Int(number))
	}

	heap.Delete(Int(15))
	heap.Delete(Int(12))

	numbers = RemoveInts(numbers, 15)
	numbers = RemoveInts(numbers, 12)
	sort.Ints(numbers)

	for _, number := range numbers {
		if Int(number) != heap.DeleteMin().(go_heaps.Integer) {
			t.Fail()
		}
	}
}

func TestBinomialHeapStringDelete(t *testing.T) {
	heap := &BinomialHeap{}

	strings := []string{"a", "ccc", "bb", "d"}
	
	for _, str := range strings {
		heap.Insert(Str(str))
	}

	heap.Delete(Str("ccc"))

	strings = RemoveStrs(strings, "ccc")
	sort.Strings(strings)

	for _, str := range strings {
		if Str(str) != heap.DeleteMin().(go_heaps.String) {
			t.Fail()
		}
	}
}

func TestLeftistHeapString(t *testing.T) {
	heap := &BinomialHeap{}

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

func RemoveInts(s []int, hay int) []int {
	sort.Ints(s)
	i := sort.SearchInts(s, hay)
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveStrs(s []string, hay string) []string {
	sort.Strings(s)
	i := sort.SearchStrings(s, hay)
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}

func Str(value string) go_heaps.String {
	return go_heaps.String(value)
}
