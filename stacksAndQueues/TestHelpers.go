package stacksAndQueues

import (
	"testing"
)


func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
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
