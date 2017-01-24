package stacksAndQueues

import (
	"testing"
)

var sampleValues = []int{100, 101, 102, 103, 104}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
