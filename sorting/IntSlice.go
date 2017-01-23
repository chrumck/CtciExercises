package sorting

import (
	"sort"
)

// IntSlice is an extended version of sort.IntSlice which implements SortInterface
type IntSlice struct {
	sort.IntSlice
}

// GetValue gets the i item of the IntSlice
func (intSlice IntSlice) GetValue(i int) interface{} {
	return intSlice.IntSlice[i]
}

// SetValue sets the i item of the IntSlice
func (intSlice IntSlice) SetValue(i int, value interface{}) {
	intSlice.IntSlice[i] = value.(int)
}
