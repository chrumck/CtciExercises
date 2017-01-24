package linkedLists

import (
	"container/list"
	"testing"
)

func TestRemoveDupes_WithExampleLists_ReturnsWithNoDupes(test *testing.T) {
	// Arrange
	exampleLists := GetExampleLists()
	for _, exampleList := range exampleLists {
		// Act
		filteredList := RemoveDupes(exampleList)
		// Assert
		if ContainsDupes(filteredList) {
			test.Error("The filtered list contains dupes but it should not.")
		}
	}
}

func TestRemoveDupes_WhenListEmpty_ReturnsEmpty(test *testing.T) {
	// Act
	filteredList := RemoveDupes(list.New())
	// Assert
	if filteredList.Len() != 0 {
		test.Error("The filtered list should be empty but is not.")
	}
}

func TestRemoveDupes2_WithExampleLists_ReturnsWithNoDupes(test *testing.T) {
	// Arrange
	exampleLists := GetExampleLists()
	for _, exampleList := range exampleLists {
		// Act
		RemoveDupes2(exampleList)
		// Assert
		if ContainsDupes(exampleList) {
			test.Error("The filtered list contains dupes but it should not.")
		}
	}
}

func TestRemoveDupes2_WhenListEmpty_ReturnsEmpty(test *testing.T) {
	// Arrange
	list := list.New()
	// Act
	RemoveDupes2(list)
	// Assert
	if list.Len() != 0 {
		test.Error("The filtered list should be empty but is not.")
	}
}
