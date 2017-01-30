package recursionAndDynamicProgramming

import (
	"testing"
)

func TestTripleStep_With0Steps_ReturnsCorrectResult(test *testing.T){

	// Arrange
	sampleTotalSteps := 0
	expectedWays := 1

	// Act
	result := TripleStep(sampleTotalSteps)

	// Assert
	if result != expectedWays {
		test.Errorf("The result should be %v, but is %v.", expectedWays, result)
	}
}


func TestTripleStep_WithOneStep_ReturnsCorrectResult(test *testing.T){

	// Arrange
	sampleTotalSteps := 1
	expectedWays := 1

	// Act
	result := TripleStep(sampleTotalSteps)

	// Assert
	if result != expectedWays {
		test.Errorf("The result should be %v, but is %v.", expectedWays, result)
	}
}

func TestTripleStep_WithTwoSteps_ReturnsCorrectResult(test *testing.T){

	// Arrange
	sampleTotalSteps := 2
	expectedWays := 2

	// Act
	result := TripleStep(sampleTotalSteps)

	// Assert
	if result != expectedWays {
		test.Errorf("The result should be %v, but is %v.", expectedWays, result)
	}
}

func TestTripleStep_WithThreeSteps_ReturnsCorrectResult(test *testing.T){

	// Arrange
	sampleTotalSteps := 3
	expectedWays := 4

	// Act
	result := TripleStep(sampleTotalSteps)

	// Assert
	if result != expectedWays {
		test.Errorf("The result should be %v, but is %v.", expectedWays, result)
	}
}

func TestTripleStep_WithFourSteps_ReturnsCorrectResult(test *testing.T){

	// Arrange
	sampleTotalSteps := 4
	expectedWays := 7

	// Act
	result := TripleStep(sampleTotalSteps)

	// Assert
	if result != expectedWays {
		test.Errorf("The result should be %v, but is %v.", expectedWays, result)
	}
}

func TestTripleStep_WithFiveSteps_ReturnsCorrectResult(test *testing.T){

	// Arrange
	sampleTotalSteps := 5
	expectedWays := 13

	// Act
	result := TripleStep(sampleTotalSteps)

	// Assert
	if result != expectedWays {
		test.Errorf("The result should be %v, but is %v.", expectedWays, result)
	}
}