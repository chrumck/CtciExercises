package sorting

import (
	"math/rand"
	"sort"
)

// Shuffle shuffles the elements in given data
func Shuffle(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		data.Swap(i, rand.Intn(i+1))
	}
}
