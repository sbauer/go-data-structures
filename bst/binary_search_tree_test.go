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
