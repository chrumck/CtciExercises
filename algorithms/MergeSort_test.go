// Exercise: implement quicksort algorithm - tests.

package algorithms

import (
	"testing"
)

func TestMergeSort_WhenSortingTestData_ReturnsSorted(test *testing.T) {
	for _, testSample := range sortTestSamples {
		MergeSort(testSample)
		if !isSorted(testSample) {
			test.Errorf("Data not in order: %v", testSample)
		}
	}
}


