package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateQueue(t *testing.T) {
	queue := NewQueue()

	assert.NotNil(t, queue)
}

func Test_EmptyShouldReturnTrueForEmptyQueue(t *testing.T) {
	queue := NewQueue()

	assert.True(t, queue.Empty())
}

func Test_EmptyShouldReturnFalseForNonEmptyQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("test")

	assert.False(t, queue.Empty())
}

func Test_EnqueueShouldAddElementToSlice(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("test")

	assert.Equal(t, 1, len(queue.data))
	assert.Equal(t, "test", queue.data[0])
}

func Test_DequeueShouldRemoveAndReturnFirstElementItemAdded(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("first")
	queue.Enqueue("second")

	assert.Equal(t, 2, queue.Length())

	item := queue.Dequeue()

	assert.Equal(t, "first", item)
	assert.Equal(t, 1, queue.Length())
}

func Test_DequeueShouldReturnNilForEmptyQueue(t *testing.T) {
	queue := NewQueue()

	assert.True(t, queue.Empty())
	assert.Nil(t, queue.Dequeue())
}

func Test_PeekShouldReturnNilForEmptyQueue(t *testing.T) {
	queue := NewQueue()

	assert.True(t, queue.Empty())
	assert.Nil(t, queue.Peek())
}

func Test_PeekShouldReturnAndKeepFirstElementItemAdded(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("first")
	queue.Enqueue("second")

	assert.Equal(t, 2, queue.Length())

	item := queue.Peek()

	assert.Equal(t, "first", item)
	assert.Equal(t, 2, queue.Length())
}

func Test_LengthShouldReflectNumberOfItems(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	assert.Equal(t, 3, queue.Length())
}
