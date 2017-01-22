// Exercise: implement quicksort algorithm - tests.

package sorting

import (
	"testing"
)

func TestQuickSort_WhenSortingTestData_ReturnsSorted(test *testing.T) {
	for _, testSample := range sortTestSamples {
		QuickSort(testSample)
		if !isSorted(testSample) {
			test.Errorf("Data not in order: %v", testSample)
		}
	}
}
