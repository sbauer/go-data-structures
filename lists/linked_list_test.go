package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateLinkedList(t *testing.T) {
	list := &LinkedList{}

	assert.NotNil(t, list)
}

func Test_InsertCreatesNewHead(t *testing.T) {
	list := &LinkedList{}

	assert.Nil(t, list.head)
	list.Insert("testing")

	assert.NotNil(t, list.head)
	assert.Equal(t, "testing", list.head.data)
}

func Test_InsertMovesPreviousHeadToNext(t *testing.T) {
	list := &LinkedList{}

	list.Insert("first insert")
	list.Insert("second insert")

	assert.Equal(t, "second insert", list.head.data)
	assert.Equal(t, "first insert", list.head.next.data)
}

func Test_AppendCreatesNewHeadWithEmptyList(t *testing.T) {
	list := &LinkedList{}

	list.Append("first")

	assert.NotNil(t, list.head)
	assert.Equal(t, "first", list.head.data)
}

func Test_AppendAddsNodeToEnd(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")
	list.Append("appended")

	assert.NotNil(t, list.head.next)
	assert.Equal(t, "appended", list.head.next.data)
}

func Test_ContainsShouldReturnFalseForEmptyList(t *testing.T) {
	list := &LinkedList{}

	assert.False(t, list.Contains("missing"))
}

func Test_ContainsShouldReturnFalseForNoMatches(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	assert.False(t, list.Contains("missing"))
}

func Test_ContainsShouldReturnTrueForHeadMatch(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	assert.True(t, list.Contains("first"))
}

func Test_ContainsShouldReturnTrueForSecondaryNodeMatch(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")
	list.Append("second")
	list.Append("third")

	assert.True(t, list.Contains("second"))
	assert.True(t, list.Contains("third"))
}

func Test_FindShouldReturnNilForEmptyList(t *testing.T) {
	list := &LinkedList{}

	assert.Nil(t, list.Find("missing"))
}

func Test_FindShouldReturnNilForNoMatches(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	assert.Nil(t, list.Find("missing"))
}

func Test_FindShouldReturnTrueForHeadMatch(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	result := list.Find("first")

	assert.NotNil(t, result)
	assert.Equal(t, "first", result.data)
}

func Test_FindShouldReturnTrueForSecondaryNodeMatch(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")
	list.Append("second")
	list.Append("third")

	secondMatch := list.Find("second")
	thirdMatch := list.Find("third")

	assert.NotNil(t, secondMatch)
	assert.Equal(t, "second", secondMatch.data)
	assert.NotNil(t, thirdMatch)
	assert.Equal(t, "third", thirdMatch.data)
}

func Test_RemoveShouldReturnFalseForEmptyList(t *testing.T) {
	list := &LinkedList{}
	node := &Node{data: "missing"}

	assert.False(t, list.Remove(node))
}

func Test_RemoveShouldReturnFalseForNilNode(t *testing.T) {
	list := &LinkedList{}

	assert.False(t, list.Remove(nil))
}

func Test_RemoveShouldReturnFalseForNoNodeMatches(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	node := &Node{data: "missing"}

	assert.False(t, list.Remove(node))
}

func Test_RemoveShouldReturnTrueAfterRemovingHeadNode(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	node := list.Find("first")

	assert.True(t, list.Remove(node))
}

func Test_RemoveShouldReturnTrueAfterRemovingSecondaryNode(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")
	list.Append("second")

	node := list.Find("second")

	assert.True(t, list.Remove(node))
}

func Test_RemoveShouldDecreaseLengthAfterRemovingNode(t *testing.T) {
	list := &LinkedList{}
	list.Insert("first")

	node := list.Find("first")

	assert.Equal(t, 1, list.length)
	assert.True(t, list.Remove(node))
	assert.Equal(t, 0, list.length)
}
