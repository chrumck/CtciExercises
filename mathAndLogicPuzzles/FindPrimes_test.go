package mathAndLogicPuzzles

import (
	"testing"
)

func TestFindPrimes_WhenValidNGiven_ReturnsPrimesCorrectly(test *testing.T)  {
	
	// Arrange
	n := 16
	expected := []int{2, 3, 5, 7, 11, 13}

	// Act
	results := FindPrimes(n)

	// Assert
	if len(results) != len(expected) {
		test.Errorf("The result: %v should be %v \n", results, expected)
	}
	for i := range expected {
		if results[i] != expected[i] {
			test.Errorf("The result: %v should be %v \n", results, expected)
		}	
	}
	
}
