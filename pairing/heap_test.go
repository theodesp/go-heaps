package pairing

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
	heap "github.com/theodesp/go-heaps"
	"fmt"
	"math/rand"
	"time"
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
	t.Do(func(a heap.Item) {
		out = append(out, a)
		return
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

func (suite *PairingHeapTestSuite) TestIsEmpty() {
	assert.Equal(suite.T(), suite.heap.IsEmpty(), true)
	suite.heap.Insert(Int(4))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(1))

	assert.Equal(suite.T(), suite.heap.IsEmpty(), false)
}

func (suite *PairingHeapTestSuite) TestFindMin() {
	suite.heap.Insert(Int(4))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(6))
	suite.heap.Insert(Int(3))

	assert.Equal(suite.T(), suite.heap.FindMin(), Int(2))
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
}

func (suite *PairingHeapTestSuite) TestFind() {
	item := suite.heap.Find(Int(10))
	assert.Nil(suite.T(), item)

	suite.heap.Insert(Int(4))

	item = suite.heap.Find(Int(4))
	assert.NotNil(suite.T(),item)
	assert.Equal(suite.T(),item, Int(4))

	suite.heap.Insert(Int(8))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(Int(5))
	suite.heap.Insert(Int(3))
	suite.heap.Insert(Int(9))

	item = suite.heap.Find(Int(9))
	assert.NotNil(suite.T(),item)
	assert.Equal(suite.T(),item, Int(9))
}

func (suite *PairingHeapTestSuite) TestAdjust() {
	for _, v := range rang(10) {
		suite.heap.Insert(v)
	}
	for i, item := range rangrev(10) {
		assert.NotNil(suite.T(), suite.heap.Adjust(item, Int(i)))
	}
	//assert.Nil(suite.T(), suite.heap.DeleteMin())
	//suite.heap.Insert(Int(4))
	//suite.heap.Insert(Int(8))
	//suite.heap.Insert(Int(2))
	//suite.heap.Insert(Int(5))
	//suite.heap.Insert(Int(3))
	//suite.heap.Insert(Int(9))
	//
	//root := Int(2)
	//assert.NotNil(suite.T(), suite.heap.Adjust(Int(2), Int(10)))
	//assert.NotEqual(suite.T(), suite.heap.FindMin(), root)
	//assert.NotNil(suite.T(), suite.heap.Find(Int(10)))
	//assert.NotNil(suite.T(), suite.heap.Find(Int(9)))
	//
	//assert.Nil(suite.T(), suite.heap.Adjust(Int(2), Int(5)))
	//assert.NotNil(suite.T(), suite.heap.Adjust(Int(10), Int(13)))
	//assert.NotNil(suite.T(), suite.heap.Adjust(Int(9), Int(5)))
	//assert.NotNil(suite.T(), suite.heap.Find(Int(13)))
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