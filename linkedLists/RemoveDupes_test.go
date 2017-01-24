package linkedLists

import (
	"container/list"
	"testing"
)

func TestRemoveDupes_WithExampleLists_ReturnsWithNoDupes(test *testing.T) {
	exampleLists := GetExampleLists()
	for _, exampleList := range exampleLists {
		filteredList := RemoveDupes(exampleList)
		if ContainsDupes(filteredList) {
			test.Error("The filtered list contains dupes but it should not.")
		}
	}
}

func TestRemoveDupes_WhenListEmpty_ReturnsEmpty(test *testing.T) {
	filteredList := RemoveDupes(list.New())
	if filteredList.Len() != 0 {
		test.Error("The filtered list should be empty but is not.")
	}
}

func TestRemoveDupes2_WithExampleLists_ReturnsWithNoDupes(test *testing.T) {
	exampleLists := GetExampleLists()
	for _, exampleList := range exampleLists {
		RemoveDupes2(exampleList)
		if ContainsDupes(exampleList) {
			test.Error("The filtered list contains dupes but it should not.")
		}
	}
}

func TestRemoveDupes2_WhenListEmpty_ReturnsEmpty(test *testing.T) {
	list := list.New()
	RemoveDupes2(list)
	if list.Len() != 0 {
		test.Error("The filtered list should be empty but is not.")
	}
}
