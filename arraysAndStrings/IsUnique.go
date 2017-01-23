package arraysAndStrings

import (
	"github.com/chrumck/CtciExercises/sorting"
)

// IsUnique checks if all characters in the given string are unique.
func IsUnique(s string) bool {
	foundRunes := map[rune]bool{}

	for _, c := range s {
		if foundRunes[c] == true {
			return false
		}
		foundRunes[c] = true
	}
	return true
}

// IsUnique2 checks if all characters in the given string are unique.
// It uses a 16 byte bit vector - it means it works for ASCII range only.
func IsUnique2(s string) bool {
	const bitsInByte = 8
	foundRunes := [16]byte{}
	for _, c := range s {
		byteIndex := c / bitsInByte
		byteMask := byte(1 << uint(c%bitsInByte))

		if foundRunes[byteIndex]&byteMask > 0 {
			return false
		}
		foundRunes[byteIndex] = foundRunes[byteIndex] | byteMask
	}
	return true
}

// IsUnique3 checks if all characters in the given string are unique.
// Sorts the string and then check if there are any duplicates. Takes O(n) time
func IsUnique3(s string) bool {
	runes := []rune(sorting.SortString(s))
	runesLength := len(runes)
	for i := 0; i < runesLength-1; i++ {
		if runes[i] == runes[i+1] {
			return false
		}
	}
	return true
}
