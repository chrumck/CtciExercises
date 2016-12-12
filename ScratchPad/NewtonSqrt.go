package ScratchPad

import (
	"fmt"
	"math"
)

const maxDelta = 0.02

// NewtonSqrt returns the square root of the number using the Newton approximation
func NewtonSqrt(x float64) (sqrt float64) {
	sqrt = 1
	deltaZ := float64(1)

	for i := 1; deltaZ > maxDelta; i++ {
		if i%100 == 0 {
			fmt.Println("i")
		}
		next := sqrt - (sqrt*sqrt-x)/2*sqrt
		deltaZ = math.Abs(sqrt - next)
		fmt.Printf("Iteration: %v, delta: %v\n", i, deltaZ)
		sqrt = next
	}

	return
}
