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

func pushAndPopValues(stack ThreeInOneStack, sampleValues []int, passCount int) bool {
	for pass := 1; pass <= passCount; pass++ {
		for _, sampleValue := range sampleValues {
			for i := 0; i <= 2; i++ {
				stack.Push(i, sampleValue)
			}
		}

		for i := len(sampleValues) - 1; i >= 0; i-- {
			for j := 0; j <= 2; j++ {
				if stack.Pop(j) != sampleValues[i] {
					return false
				}
			}
		}
	}
	return true
}
