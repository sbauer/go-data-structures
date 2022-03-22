package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateStack(t *testing.T) {
	stack := NewStack()

	assert.NotNil(t, stack)
}

func Test_EmptyShouldBeTrueForNewStack(t *testing.T) {
	stack := NewStack()

	assert.True(t, stack.Empty())
}

func Test_LengthShouldBeZeroForNewStack(t *testing.T) {
	stack := NewStack()

	assert.Equal(t, 0, stack.Length())
}

func Test_LengthShouldBeAccurateAfterPush(t *testing.T) {
	stack := NewStack()

	stack.Push("test")
	assert.Equal(t, 1, stack.Length())
}

func Test_PopShouldReturnMostRecentPushAndRemoveIt(t *testing.T) {
	stack := NewStack()

	stack.Push("test")
	stack.Push("second")

	result := stack.Pop()

	assert.Equal(t, "second", result)
	assert.Equal(t, 1, stack.Length())

	secondResult := stack.Pop()

	assert.Equal(t, "test", secondResult)
	assert.Equal(t, 0, stack.Length())
}

func Test_PopShouldReturnNilIfStackIsEmpty(t *testing.T) {
	stack := NewStack()

	assert.Nil(t, stack.Pop())
}
