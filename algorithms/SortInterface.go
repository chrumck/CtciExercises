package algorithms

import (
	"sort"
)

// SortInterface extends sort.Interface with methods allowing accessing items of the collection
type SortInterface interface {
	sort.Interface
	GetValue(i int) interface{}
	SetValue(i int, value interface{})
}
