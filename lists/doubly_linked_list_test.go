package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateDoublyLinkedList(t *testing.T) {
	list := &DoublyLinkedList{}

	assert.NotNil(t, list)
}

func Test_EmptyShouldBeTrueForEmptyList(t *testing.T) {
	list := &DoublyLinkedList{}

	assert.True(t, list.Empty())
}

func Test_InsertShouldCreateNewHeadAndTailIfEmpty(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Insert("first")

	assert.NotNil(t, list.head)
	assert.Equal(t, "first", list.head.data)
	assert.Equal(t, list.head, list.tail)
}

func Test_InsertShouldCreateNewHeadAndKeepTail(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Insert("first")
	list.Insert("second")

	assert.NotNil(t, list.head)
	assert.Equal(t, "second", list.head.data)
	assert.Equal(t, "first", list.tail.data)
}

func Test_InsertShouldIncrementLength(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Insert("first")
	list.Insert("second")

	assert.Equal(t, 2, list.length)
}

func Test_AppendShouldCreateNewHeadAndTailIfEmpty(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append("first")

	assert.NotNil(t, list.head)
	assert.Equal(t, "first", list.head.data)
	assert.Equal(t, list.head, list.tail)
}

func Test_AppendShouldKeepHeadAndUpdateTail(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Insert("first")
	list.Append("second")

	assert.Equal(t, "second", list.tail.data)
}

func Test_DoublyLinkedListAppendShouldIncrementLength(t *testing.T) {
	list := &DoublyLinkedList{}
	list.Append("first")
	list.Append("second")

	assert.Equal(t, 2, list.length)
}

func Test_DoublyLinkedListContainsShouldReturnFalseForEmptyList(t *testing.T) {
	list := &DoublyLinkedList{}

	assert.False(t, list.Contains("nothing"))
}

func Test_DoublyLinkedListContainsShouldReturnTrueForHeadMatch(t *testing.T) {
	list := &DoublyLinkedList{}

	list.Insert("first")
	assert.True(t, list.Contains("first"))
}

func Test_DoublyLinkedListContainsShouldReturnTrueForOtherMatch(t *testing.T) {
	list := &DoublyLinkedList{}

	list.Insert("first")
	list.Append("second")
	list.Append("third")
	assert.True(t, list.Contains("second"))
}

func Test_DoublyLinkedListFindShouldReturnNilForEmptyList(t *testing.T) {
	list := &DoublyLinkedList{}

	assert.Nil(t, list.Find("nothing"))
}

func Test_DoublyLinkedListFindShouldReturnNodeForHeadMatch(t *testing.T) {
	list := &DoublyLinkedList{}

	list.Insert("first")
	result := list.Find("first")

	assert.NotNil(t, result)
	assert.Equal(t, "first", result.data)
}

func Test_DoublyLinkedListFindShouldReturnNodeForOtherMatch(t *testing.T) {
	list := &DoublyLinkedList{}

	list.Insert("first")
	list.Append("second")
	list.Append("third")
	result := list.Find("second")

	assert.NotNil(t, result)
	assert.Equal(t, "second", result.data)
}
