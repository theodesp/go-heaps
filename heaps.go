package go_heaps

// Interface is basic interface that all Heaps implement.
type Interface interface {
	// Inserts an element to the heap and returns it
	Insert(v interface{}) interface{}

	// DeleteMin deletes and returns the smallest element
	DeleteMin() interface{}

	// FindMin returns the minimum element
	FindMin() interface{}

	// Removes all items
	Clear()
}

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator func(a, b interface{}) int

// StringComparator provides a fast comparison on strings
func StringComparator(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// IntComparator provides a basic comparison on int
func IntComparator(a, b interface{}) int {
	a1 := a.(int)
	a2 := b.(int)
	switch {
	case a1 > a2:
		return 1
	case a1 < a2:
		return -1
	default:
		return 0
	}
}
