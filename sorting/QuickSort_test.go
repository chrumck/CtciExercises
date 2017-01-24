package sorting

import (
	"testing"
)

func TestQuickSort_WhenSortingTestData_ReturnsSorted(test *testing.T) {
	// Arrange
	for _, testSample := range sortTestSamples {
		// Act
		QuickSort(testSample)
		// Assert
		if !isSorted(testSample) {
			test.Errorf("Data not in order: %v", testSample)
		}
	}
}
