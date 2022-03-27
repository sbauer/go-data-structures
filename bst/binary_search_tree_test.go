package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCreateTree(t *testing.T) {
	assert.NotNil(t, NewBinarySearchTree())
}

func Test_InsertSingleEntryExists(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)

	assert.NotNil(t, tree.root)
	assert.Equal(t, 5, tree.root.value)
}

func Test_InsertPutsSmallerEntryToLeft(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)

	assert.Nil(t, tree.root.right)
	assert.NotNil(t, tree.root.left)
	assert.Equal(t, 3, tree.root.left.value)
}

func Test_InsertPutsLargerEntryToRight(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(10)

	assert.Nil(t, tree.root.left)
	assert.NotNil(t, tree.root.right)
	assert.Equal(t, 10, tree.root.right.value)
}

func Test_FindReturnsNilForEmptyTree(t *testing.T) {
	tree := NewBinarySearchTree()

	assert.Nil(t, tree.Find(2))
}

func Test_FindReturnsNodeAtRoot(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)

	result := tree.Find(5)
	assert.NotNil(t, result)
	assert.Equal(t, 5, result.value)
}

func Test_FindCanLocateInOtherLeaves(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(15)

	result := tree.Find(1)
	secondResult := tree.Find(15)
	assert.NotNil(t, result)
	assert.NotNil(t, secondResult)
	assert.Equal(t, 1, result.value)
	assert.Equal(t, 15, secondResult.value)
}

func Test_ContainsReturnsFalseForEmptyTree(t *testing.T) {
	tree := NewBinarySearchTree()

	assert.False(t, tree.Contains(5))
}

func Test_ContainsReturnsTrueForMatches(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)

	assert.True(t, tree.Contains(5))
	assert.True(t, tree.Contains(3))
	assert.True(t, tree.Contains(1))
}

func Test_ContainsReturnsFalseForNoMatches(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)

	assert.False(t, tree.Contains(6))
	assert.False(t, tree.Contains(2))
	assert.False(t, tree.Contains(11))
}

func Test_MinimumReturnsErrorForEmptyTree(t *testing.T) {
	tree := NewBinarySearchTree()

	minimum, err := tree.Minimum()

	assert.Equal(t, -1, minimum)
	assert.NotNil(t, err)
}

func Test_MinimumReturnsSingleRoot(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)

	minimum, err := tree.Minimum()

	assert.Equal(t, 5, minimum)
	assert.Nil(t, err)
}

func Test_MinimumReturnsCorrectMinimum(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(-3)
	tree.Insert(100)

	minimum, err := tree.Minimum()

	assert.Equal(t, -3, minimum)
	assert.Nil(t, err)
}

func Test_MaximumReturnsErrorForEmptyTree(t *testing.T) {
	tree := NewBinarySearchTree()

	max, err := tree.Maximum()

	assert.Equal(t, -1, max)
	assert.NotNil(t, err)
}

func Test_MaximumReturnsSingleRoot(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)

	max, err := tree.Maximum()

	assert.Equal(t, 5, max)
	assert.Nil(t, err)
}

func Test_MaximumReturnsCorrectMinimum(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(150)
	tree.Insert(-3)
	tree.Insert(100)

	max, err := tree.Maximum()

	assert.Equal(t, 150, max)
	assert.Nil(t, err)
}

func Test_DeleteOnEmptyTrueDoesntPanic(t *testing.T) {
	tree := NewBinarySearchTree()
	assert.NotPanics(t, func() {
		tree.Delete(5)
	})
}

func Test_DeleteShouldDeleteRoot(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)

	assert.NotNil(t, tree.root)

	tree.Delete(5)

	assert.Nil(t, tree.root)
}

func Test_DeleteShouldRemoveLeafWithoutChildren(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(10)

	assert.Equal(t, 3, tree.root.left.value)
	tree.Delete(3)
	assert.Nil(t, tree.root.left)
}

func Test_DeleteShouldReplaceWithRightChildIfLeftDoesntExist(t *testing.T) {
	/*
				5
			  /   \
			 3     10
			/ \    / \
		              15
	*/
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(15)

	assert.Equal(t, 10, tree.root.right.value)
	assert.Equal(t, 15, tree.root.right.right.value)

	tree.Delete(10)

	assert.Equal(t, 15, tree.root.right.value)
	assert.Nil(t, tree.root.right.right)
}

func Test_DeleteShouldReplaceWithLeftChildIfRightDoesntExist(t *testing.T) {
	/*
				5
			  /   \
			 3     10
			/ \    / \
		          9
	*/
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(9)

	assert.Equal(t, 10, tree.root.right.value)
	assert.Equal(t, 9, tree.root.right.left.value)

	tree.Delete(10)

	assert.Equal(t, 9, tree.root.right.value)
	assert.Nil(t, tree.root.right.left)
}

func Test_DeleteShouldReplaceWithRightChildIfBothExist(t *testing.T) {
	/*
				5
			  /   \
			 3     10
			/ \    / \
		          9  11
	*/
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(11)

	assert.Equal(t, 10, tree.root.right.value)
	assert.Equal(t, 9, tree.root.right.left.value)
	assert.Equal(t, 11, tree.root.right.right.value)

	tree.Delete(10)

	assert.Equal(t, 11, tree.root.right.value)
	assert.Nil(t, tree.root.right.right)
	assert.Equal(t, 9, tree.root.right.left.value)
}

func Test_DeleteShouldReplaceWithSmallestRightNodeIfBothChildrenExistWithChildren(t *testing.T) {
	/*
						5
					  /   \
					 3     10
					/ \    / \
				          9  15
						 /  /  \
		                7  13   18
						    \
							 14
	*/
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(18)
	tree.Insert(14)

	assert.Equal(t, 10, tree.root.right.value)

	tree.Delete(10)

	assert.Equal(t, 13, tree.root.right.value)
	assert.Equal(t, 14, tree.root.right.right.left.value)
}

func Test_DeleteShouldReplaceWithRootWithCorrectSmallestRightChild(t *testing.T) {
	/*
						5
					  /   \
					 3     10
					/ \    / \
				          9  15
						 /  /  \
		                7  13   18
						    \
							 14
	*/
	tree := NewBinarySearchTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(18)
	tree.Insert(14)

	assert.Equal(t, 5, tree.root.value)

	tree.Delete(5)

	assert.Equal(t, 7, tree.root.value)
	assert.Nil(t, tree.root.right.left.left)
}
