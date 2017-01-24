package arraysAndStrings

import (
	"testing"
)

func TestIsUnique_WhenStringNotUnique_returnsFalse(test *testing.T) {
	// Act, Assert
	if IsUnique(exampleString) {
		test.Errorf("Returned IsUnique == true but the text is not: %v", exampleString)
	}
}

func TestIsUnique_WhenStringUnique_returnsTrue(test *testing.T) {
	// Act, Assert
	if !IsUnique(exampleUniqueString) {
		test.Errorf("Returned IsUnique == false but the text is unique: %v", exampleUniqueString)
	}
}

func TestIsUnique2_WhenStringNotUnique_returnsFalse(test *testing.T) {
	// Act, Assert
	if IsUnique2(exampleString) {
		test.Errorf("Returned IsUnique == true but the text is not: %v", exampleString)
	}
}

func TestIsUnique2_WhenStringUnique_returnsTrue(test *testing.T) {
	// Act, Assert
	if !IsUnique2(exampleUniqueString) {
		test.Errorf("Returned IsUnique == false but the text is unique: %v", exampleUniqueString)
	}
}

func TestIsUnique3_WhenStringNotUnique_returnsFalse(test *testing.T) {
	// Act, Assert
	if IsUnique3(exampleString) {
		test.Errorf("Returned IsUnique == true but the text is not: %v", exampleString)
	}
}

func TestIsUnique3_WhenStringUnique_returnsTrue(test *testing.T) {
	// Act, Assert
	if !IsUnique3(exampleUniqueString) {
		test.Errorf("Returned IsUnique == false but the text is unique: %v", exampleUniqueString)
	}
}
