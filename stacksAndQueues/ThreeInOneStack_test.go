package stacksAndQueues

import (
	"testing"
)

func TestPushPop_WhenValuesPushedAndPopped_PushesAndPopsCorrectly(test *testing.T) {
	stack := ThreeInOneStack{}
	for _, sampleValue := range sampleValues {
		for i := 0; i <= 2; i++ {
			stack.Push(i, sampleValue)
		}
	}

	for i := len(sampleValues) - 1; i < 0; i-- {
		for j := 0; j <= 2; j++ {
			if stack.Pop(j) != sampleValues[i] {
				test.Error("The filtered list should be empty but is not.")
			}
		}
	}
}

func TestPushPop_WhenEmptyStackPopped_Panics(test *testing.T) {
	stack := ThreeInOneStack{}
	assertPanic(test, func() { stack.Pop(0) })
}
