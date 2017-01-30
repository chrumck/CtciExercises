package recursionAndDynamicProgramming

// TripleStep calculates in how many ways a child can hop the stairs if it takes 1, 2, or 3 steps.
func TripleStep(n int) (ways int) {
	const maxSteps = 3

	var tripleStep func(stepCountSoFar int)
	tripleStep = func(stepCountSoFar int) {

		if stepCountSoFar == n {
			ways++
			return
		}

		for i := 1; i <= maxSteps; i++ {
			if i <= n-stepCountSoFar {
				tripleStep(stepCountSoFar + i)
			}
		}
		return
	}

	tripleStep(0)
	return
}
