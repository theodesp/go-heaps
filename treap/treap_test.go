package treap

import (
	"sort"
	"testing"

	goheap "github.com/theodesp/go-heaps"
)

func TestTreapInteger(t *testing.T) {
	treap := New()

	numbers := []int{4, 3, 2, 5}

	for _, number := range numbers {
		treap.Insert(goheap.Integer(number))
	}

	sort.Ints(numbers)

	for _, number := range numbers {
		if goheap.Integer(number) != treap.DeleteMin().(goheap.Integer) {
			t.Fail()
		}
	}
}

func TestTreapString(t *testing.T) {
	treap := New()

	strs := []string{"a", "ccc", "bb", "d"}

	for _, str := range strs {
		treap.Insert(goheap.String(str))
	}

	sort.Strings(strs)

	for _, str := range strs {
		if goheap.String(str) != treap.DeleteMin().(goheap.String) {
			t.Fail()
		}
	}
}
