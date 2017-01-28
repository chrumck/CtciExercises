package sorting

// Comparer interface represents any types which can be compared
type Comparer interface {
	CompareTo(other Comparer) int
}

