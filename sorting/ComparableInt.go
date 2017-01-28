package sorting

// ComparableInt represents an int with a method CompareTo fulfilling the Comparer interface.
type ComparableInt int

// CompareTo compares to other ComparableInt and returns -1, 0, or 1
func (cInt ComparableInt) CompareTo(other Comparer) int {
	otherT := other.(ComparableInt)
	switch {
	case cInt < otherT:
		return -1
	case cInt > otherT:
		return 1
	default:
		return 0
	}
}

