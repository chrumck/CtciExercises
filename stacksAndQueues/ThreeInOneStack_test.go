package stacksAndQueues

import (
	"testing"
)

func TestPushPop_WhenValuesPushedAndPopped_PushesAndPopsCorrectly(test *testing.T) {
	
	// Arrange
	stack := ThreeInOneStack{}
	var sampleValues = []int{100, 101, 102, 103, 104}
	
	// Act
	passed := pushAndPopValues(stack, sampleValues, 2)
	
	// Assert
	if !passed {
		test.Error("The filtered list should be empty but is not.")
	}
}

func TestPop_WhenEmptyStackPopped_Panics(test *testing.T) {
	
	// Arrange
	stack := ThreeInOneStack{}
	
	// Act, Assert
	assertPanic(test, func() { stack.Pop(0) })
}

