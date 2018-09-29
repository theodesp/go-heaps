package pairing

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
)

type PairingHeapTestSuite struct {
	suite.Suite
	heap *PairHeap
}

func (suite *PairingHeapTestSuite) SetupTest() {
	suite.heap = NewWithIntComparator()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(PairingHeapTestSuite))
}

func (suite *PairingHeapTestSuite) TestIsEmpty() {
	assert.Equal(suite.T(), suite.heap.IsEmpty(), true)
	suite.heap.Insert(4)
	suite.heap.Insert(2)
	suite.heap.Insert(1)

	assert.Equal(suite.T(), suite.heap.IsEmpty(), false)
}

func (suite *PairingHeapTestSuite) TestFindMin() {
	suite.heap.Insert(4)
	suite.heap.Insert(2)
	suite.heap.Insert(6)
	suite.heap.Insert(3)

	assert.Equal(suite.T(), suite.heap.FindMin(), 2)
}

func (suite *PairingHeapTestSuite) TestDeleteMin() {
	suite.heap.Insert(4)
	suite.heap.Insert(8)
	suite.heap.Insert(6)
	suite.heap.Insert(3)

	assert.Equal(suite.T(), suite.heap.DeleteMin(), 3)
	assert.Equal(suite.T(), suite.heap.DeleteMin(), 4)
	assert.Equal(suite.T(), suite.heap.DeleteMin(), 6)
	assert.Equal(suite.T(), suite.heap.DeleteMin(), 8)
	assert.Nil(suite.T(), suite.heap.DeleteMin())
}

func (suite *PairingHeapTestSuite) TestInsert() {
	n1 := suite.heap.Insert(4)
	assert.Equal(suite.T(), n1, suite.heap.FindMin())

	n2 := suite.heap.Insert(6)
	assert.NotEqual(suite.T(), n2, suite.heap.FindMin())

	n3 := suite.heap.DeleteMin()
	assert.NotEqual(suite.T(), n3, suite.heap.FindMin())
}

func (suite *PairingHeapTestSuite) TestFind() {
	node := suite.heap.Find(10)
	assert.Nil(suite.T(), node)

	suite.heap.Insert(4)

	node = suite.heap.Find(4)
	assert.NotNil(suite.T(),node)
	assert.Equal(suite.T(),node.Value, 4)

	suite.heap.Insert(8)
	suite.heap.Insert(2)
	suite.heap.Insert(5)
	suite.heap.Insert(3)
	suite.heap.Insert(9)

	node = suite.heap.Find(9)
	assert.NotNil(suite.T(),node)
	assert.Equal(suite.T(),node.Value, 9)
}
