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
