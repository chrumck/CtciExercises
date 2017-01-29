package bitManipulation

import (
	"testing"
)

func TestInsertMIntoN(test *testing.T) {
	
	// Arrange
	M := int32(3)
	N := int32(0xFFFF)

	// Act
	result := InsertMIntoN(M, N, 2, 4)

	// Assert
	if result != 65519 {
		test.Error("Did not properly insert M into N.")
	}
}
