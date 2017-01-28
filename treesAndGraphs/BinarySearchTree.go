package treesAndGraphs

import (
	"github.com/chrumck/CtciExercises/sorting"
)

// BinarySearchTree implements the unbalanced binary search tree algorithm.
type BinarySearchTree struct {
	key   sorting.Comparer
	value interface{}
	left  *BinarySearchTree
	right *BinarySearchTree
}

// Insert walks the tree and inserts the value in the appropriate node.
func (bst *BinarySearchTree) Insert(key sorting.Comparer, value interface{}) {
	switch {
	case bst.key == nil:
		bst.key = key
		bst.value = value
		bst.left = &BinarySearchTree{}
		bst.right = &BinarySearchTree{}
	case key.CompareTo(bst.key) < 0:
		bst.left.Insert(key, value)
	case key.CompareTo(bst.key) > 0:
		bst.right.Insert(key, value)
	default:
		bst.value = value
	}
}

// Get walks the tree nodes and finds the value by the given key.
func (bst *BinarySearchTree) Get(key sorting.Comparer) interface{} {
	switch {
	case bst.key == nil:
		return nil
	case key.CompareTo(bst.key) < 0:
		return bst.left.Get(key)
	case key.CompareTo(bst.key) > 0:
		return bst.right.Get(key)
	default:
		return bst.value
	}
}

// NewFromOrderedSlice convert a slice to BST.
func NewFromOrderedSlice(values []sorting.Comparer) (bst *BinarySearchTree) {

	var convertFromSlice func(lo, hi int)
	convertFromSlice = func(lo, hi int) {
		i := (lo + hi) / 2
		bst.Insert(values[i], values[i])
		if lo < i {
			convertFromSlice(lo, i-1)
		}
		if i < hi {
			convertFromSlice(i+1, hi)
		}
	}

	hi := len(values) - 1
	if hi < 0 {
		return
	}
	bst = &BinarySearchTree{}
	convertFromSlice(0, hi)
	return
}
