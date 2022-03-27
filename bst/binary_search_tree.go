package bst

import "errors"

type Node struct {
	left  *Node
	right *Node
	value int
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (tree *BinarySearchTree) Insert(value int) {
	tree.root = tree.insert(tree.root, value)
}

func (tree *BinarySearchTree) insert(target *Node, value int) *Node {
	if target == nil {
		target = &Node{value: value}
		return target
	}

	if value < target.value {
		target.left = tree.insert(target.left, value)
	} else if value > target.value {
		target.right = tree.insert(target.right, value)
	}

	return target
}

func (tree *BinarySearchTree) Find(value int) *Node {
	return tree.find(tree.root, value)
}

func (tree *BinarySearchTree) Contains(value int) bool {
	result := tree.find(tree.root, value)

	return result != nil
}

func (tree *BinarySearchTree) find(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.value {
		return tree.find(node.left, value)
	} else if value > node.value {
		return tree.find(node.right, value)
	}

	return node
}

func (tree *BinarySearchTree) Minimum() (int, error) {
	if tree.root == nil {
		return -1, errors.New("Empty tree")
	}

	return tree.minimum(tree.root), nil
}

func (tree *BinarySearchTree) minimum(node *Node) int {
	if node.left == nil {
		return node.value
	}

	return tree.minimum(node.left)
}

func (tree *BinarySearchTree) Maximum() (int, error) {
	if tree.root == nil {
		return -1, errors.New("Empty tree")
	}

	return tree.maximum(tree.root), nil
}

func (tree *BinarySearchTree) maximum(node *Node) int {
	if node.right == nil {
		return node.value
	}

	return tree.maximum(node.right)
}

func (tree *BinarySearchTree) Delete(value int) {
	tree.root = tree.delete(tree.root, value)
}

func (tree *BinarySearchTree) delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = tree.delete(node.left, value)
	} else if value > node.value {
		node.right = tree.delete(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		node.value = tree.minimum(node.right)
		node.right = tree.delete(node.right, node.value)
	}

	return node
}
