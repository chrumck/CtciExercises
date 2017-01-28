package treesAndGraphs

import (
	"github.com/chrumck/CtciExercises/sorting"
	"testing"
)

func TestInsertAndGet_WithOneItem_InsertsAndGetsCorrectly(test *testing.T) {

	// Arrange
	var item sorting.ComparableInt = 10
	bst := &BinarySearchTree{}

	// Act
	bst.Insert(item, item)
	result := bst.Get(item)

	// Assert
	if result != item {
		test.Error("Should get the inserted item but it did not.")
	}
}

func TestInsertAndGet_WithThreeItems_InsertsAndGetsCorrectly(test *testing.T) {

	// Arrange
	bst := &BinarySearchTree{}

	// Act
	bst.Insert(sorting.ComparableInt(10), 100)
	bst.Insert(sorting.ComparableInt(11), 200)
	bst.Insert(sorting.ComparableInt(8), 300)
	result := bst.Get(sorting.ComparableInt(8))

	// Assert
	if result != 300 {
		test.Error("Should get the inserted item but it did not.")
	}
}

func TestFromOrderedSlice_WithOrderedSlice_CreatesBstCorrectly(test *testing.T) {

	// Arrange
	sampleSlice := []sorting.Comparer{
		sorting.ComparableInt(1),
		sorting.ComparableInt(2),
		sorting.ComparableInt(3),
		sorting.ComparableInt(4),
		sorting.ComparableInt(5),
		sorting.ComparableInt(6),
		sorting.ComparableInt(7),
		sorting.ComparableInt(8),
	}

	// Act
	result := NewFromOrderedSlice(sampleSlice)

	// Assert
	if result.left.left.value != sorting.ComparableInt(1) ||
		result.right.left.value != sorting.ComparableInt(5) ||
		result.right.right.right.value != sorting.ComparableInt(8){
		test.Errorf("The tree value:%v does not have expected values.", result.left.left.value)
	}
}
