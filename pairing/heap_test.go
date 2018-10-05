package pairing

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
	"github.com/mb-14/go-heaps"
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
	suite.heap.Insert(Int(4))
	suite.heap.Insert(Int(8))
	suite.heap.Insert(Int(6))
	suite.heap.Insert(Int(3))

	assert.Equal(suite.T(), suite.heap.DeleteMin(), 3)
	assert.Equal(suite.T(), suite.heap.DeleteMin(), 4)
	assert.Equal(suite.T(), suite.heap.DeleteMin(), 6)
	assert.Equal(suite.T(), suite.heap.DeleteMin(), 8)
	assert.Nil(suite.T(), suite.heap.DeleteMin())
}

func (suite *PairingHeapTestSuite) TestInsert() {
	n1 := suite.heap.Insert(Int(4))
	assert.Equal(suite.T(), n1.Value, suite.heap.FindMin())

	n2 := suite.heap.Insert(Int(6))
	assert.NotEqual(suite.T(), n2.Value, suite.heap.FindMin())

	n3 := suite.heap.DeleteMin()
	assert.NotEqual(suite.T(), n3, suite.heap.FindMin())
}

func (suite *PairingHeapTestSuite) TestFind() {
	node := suite.heap.Find(Int(10))
	assert.Nil(suite.T(), node)

	suite.heap.Insert(Int(4))

	node = suite.heap.Find(Int(4))
	assert.NotNil(suite.T(),node)
	assert.Equal(suite.T(),node.Value, Int(4))

	suite.heap.Insert(Int(8))
	suite.heap.Insert(Int(2))
	suite.heap.Insert(5)
	suite.heap.Insert(3)
	suite.heap.Insert(9)

	node = suite.heap.Find(9)
	assert.NotNil(suite.T(),node)
	assert.Equal(suite.T(),node.Value, 9)
}

func (suite *PairingHeapTestSuite) TestAdjust() {
	suite.heap.Insert(4)
	suite.heap.Insert(8)
	suite.heap.Insert(2)
	suite.heap.Insert(5)
	suite.heap.Insert(3)
	suite.heap.Insert(9)

	root := suite.heap.Root
	assert.NotNil(suite.T(), suite.heap.Adjust(root, 10))
	assert.NotEqual(suite.T(), suite.heap.FindMin(), root)
	assert.NotNil(suite.T(), suite.heap.Find(10))
	assert.NotNil(suite.T(), suite.heap.Find(9))

	assert.Nil(suite.T(), suite.heap.Adjust(suite.heap.Find(2), 5))
	assert.NotNil(suite.T(), suite.heap.Adjust(suite.heap.Find(10), 13))
	assert.NotNil(suite.T(), suite.heap.Adjust(suite.heap.Find(9), 5))
	assert.NotNil(suite.T(), suite.heap.Find(13))
}

func (suite *PairingHeapTestSuite) TestDelete() {
	suite.heap.Insert(4)
	suite.heap.Insert(8)
	suite.heap.Insert(2)
	suite.heap.Insert(5)
	suite.heap.Insert(3)
	suite.heap.Insert(9)

	assert.Nil(suite.T(), suite.heap.Delete(suite.heap.Find(10)))
	assert.NotNil(suite.T(), suite.heap.Delete(suite.heap.Find(4)))
	assert.Nil(suite.T(), suite.heap.Find(4))
	assert.NotNil(suite.T(), suite.heap.Find(8))
	assert.NotNil(suite.T(), suite.heap.Find(5))
	assert.NotNil(suite.T(), suite.heap.Find(3))
	assert.NotNil(suite.T(), suite.heap.Find(9))
}

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer{value}
}