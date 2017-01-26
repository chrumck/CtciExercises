package sorting

import (
	"testing"
)

func TestMergeSort_WhenSortingTestData_ReturnsSorted(test *testing.T) {
	
	// Arrange
	for _, testSample := range sortTestSamples {
		
		// Act
		MergeSort(testSample)
		
		// Assert
		if !isSorted(testSample) {
			test.Errorf("Data not in order: %v", testSample)
		}
	}
}


