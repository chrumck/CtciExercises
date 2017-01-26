package moderate

import (
	"testing"
)

func TestSwapNumber_WithValidParams_SwapsCorrectly(test *testing.T) {

	// Arrange
	one, two := 9, 4

	// Act
	SwapNumber(&one, &two)

	// Assert
	if one != 4 || two != 9 {
		test.Error("Did not swap correctly.")
	}
}
