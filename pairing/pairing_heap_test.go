package pairing

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	heap "github.com/theodesp/go-heaps"
)

type PairingHeapTestSuite struct {
	suite.Suite
	heap *PairHeap
}

func (suite *PairingHeapTestSuite) SetupTest() {
	suite.heap = New()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(PairingHeapTestSuite))
}

func init() {
	seed := time.Now().Unix()
	fmt.Println(seed)
	rand.Seed(seed)
}

// perm returns a random permutation of n Int items in the range [0, n).
func perm(n int) (out []heap.Item) {
	for _, v := range rand.Perm(n) {
		out = append(out, Int(v))
	}
	return
}

// rang returns an ordered list of Int items in the range [0, n).
func rang(n int) (out []heap.Item) {
	for i := 0; i < n; i++ {
		out = append(out, Int(i))
	}
	return
}

// all extracts all items from a tree in order as a slice.
func all(t *PairHeap) (out []heap.Item) {
	t.Do(func(a heap.Item) bool {
		out = append(out, a)
		return true
	})
	return
}

// rangerev returns a reversed ordered list of Int items in the range [0, n).
func rangrev(n int) (out []heap.Item) {
	for i := n - 1; i >= 0; i-- {
		out = append(out, Int(i))
	}
	return
}

func testMinHeapInvariance(suite *PairingHeapTestSuite) {
	suite.T().Helper()
	var items []heap.Item
	for {
		item := suite.heap.DeleteMin()
		if item == nil {
			break
		} else {
			items = append(items, item)
		}
	}

	for i := 0; i < len(items)-1; i += 1 {
		assert.True(suite.T(), items[i].Compare(items[i+1]) < 0)
	}
}

func (suite *PairingHeapTestSuite) TestIsEmpty() {
	assert.Equal(suite.T(), suite.heap.IsEmpty(), true)
	suite.heap.Insert(Int(4))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(1))

	assert.Equal(suite.T(), suite.heap.IsEmpty(), false)
}

func (suite *PairingHeapTestSuite) TestMeld() {
	assert.NotNil(suite.T(), suite.heap.Meld(nil))

	heapB := New()
	heapB.Insert(Int(4))
	heapB.Insert(Int(2))
	heapB.Insert(Int(6))
	heapB.Insert(Int(7))

	suite.heap.Meld(heapB)
	testMinHeapInvariance(suite)

	suite.heap.Insert(Int(4))
	suite.heap.Insert(Int(6))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(10))

	heapB = New()
	heapB.Insert(Int(1))
	heapB.Insert(Int(3))
	heapB.Insert(Int(5))
	heapB.Insert(Int(7))

	suite.heap.Meld(heapB)
	testMinHeapInvariance(suite)
}

func (suite *PairingHeapTestSuite) TestFindMin() {
	suite.heap.Insert(Int(4))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(6))
	suite.heap.Insert(Int(3))

	assert.Equal(suite.T(), suite.heap.FindMin(), Int(2))
	testMinHeapInvariance(suite)
}

func (suite *PairingHeapTestSuite) TestDeleteMin() {
	for _, v := range perm(100) {
		suite.heap.Insert(v)
	}
	var got []heap.Item
	for v := suite.heap.DeleteMin(); v != nil; v = suite.heap.DeleteMin() {
		got = append(got, v)
	}
	assert.ElementsMatch(suite.T(), got, rang(100))
}

func (suite *PairingHeapTestSuite) TestInsert() {
	for _, item := range perm(100) {
		suite.heap.Insert(item)
	}
	min := suite.heap.FindMin()
	assert.Equal(suite.T(), min, Int(0))

	got := all(suite.heap)
	want := rang(100)
	assert.ElementsMatch(suite.T(), got, want)
	testMinHeapInvariance(suite)
}

func (suite *PairingHeapTestSuite) TestFind() {
	item := suite.heap.Find(Int(10))
	assert.Nil(suite.T(), item)

	suite.heap.Insert(Int(4))

	item = suite.heap.Find(Int(4))
	assert.NotNil(suite.T(), item)
	assert.Equal(suite.T(), item, Int(4))

	suite.heap.Insert(Int(8))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(5))
	suite.heap.Insert(Int(3))
	suite.heap.Insert(Int(9))

	item = suite.heap.Find(Int(9))
	assert.NotNil(suite.T(), item)
	assert.Equal(suite.T(), item, Int(9))
	testMinHeapInvariance(suite)
}

func (suite *PairingHeapTestSuite) TestAdjust() {
	for _, v := range rang(10) {
		suite.heap.Insert(v)
	}
	for i, item := range rangrev(10) {
		assert.NotNil(suite.T(), suite.heap.Adjust(item, Int(i)))
	}
	testMinHeapInvariance(suite)
}

func (suite *PairingHeapTestSuite) TestDeleteBug() {
	vs := []int{3, 4, 5, 6, 7, 8}
	for _, v := range vs {
		suite.heap.Insert(Int(v))
	}
	assert.Equal(suite.T(), Int(3), suite.heap.Delete(Int(3)))
	assert.Nil(suite.T(), suite.heap.Find(Int(3)))
	i := suite.heap.DeleteMin()
	assert.Equal(suite.T(), Int(4), i)

	assert.Equal(suite.T(), Int(6), suite.heap.Delete(Int(6)))
	assert.Nil(suite.T(), suite.heap.Find(Int(6)))
}

func (suite *PairingHeapTestSuite) TestDelete() {
	for _, v := range rang(10) {
		suite.heap.Insert(v)
	}
	for _, item := range rangrev(10) {
		assert.NotNil(suite.T(), suite.heap.Delete(item))
	}

	assert.Nil(suite.T(), suite.heap.DeleteMin())
}

func Int(value int) heap.Integer {
	return heap.Integer(value)
}
