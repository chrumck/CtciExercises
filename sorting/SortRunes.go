package sorting

// SortRunes implements the sort.Interface and SortInterface
type SortRunes []rune

// Len returns the length of the SortRunes slice.
func (runes SortRunes) Len() int {
	return len(runes)
}

// Less indicates if if rune i is smaller than rune j.
func (runes SortRunes) Less(i, j int) bool {
	return runes[i] < runes[j]
}

// Swap swaps the rune i with the rune j.
func (runes SortRunes) Swap(i, j int) {
	runes[i], runes[j] = runes[j], runes[i]
}

// GetValue returns the value of the rune i.
func (runes SortRunes) GetValue(i int) interface{} {
	return runes[i]
}

// SetValue sets value of the rune i
func (runes SortRunes) SetValue(i int, value interface{}) {
	runes[i] = value.(rune)
}

// SortString sorts the string runes in increasing order, ordered by codepoint value.
// Uses Mergesort to sort the runes.
func SortString(s string) string {
	sortRunes := SortRunes([]rune(s))
	MergeSort(sortRunes)
	return string(sortRunes)
}
