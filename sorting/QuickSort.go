package sorting

import (
	"sort"
)

// QuickSort sorts given data using QuickSort algorithm
func QuickSort(data sort.Interface) {
	Shuffle(data) // needs to be shuffled to avoid worst case scenario
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data sort.Interface, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(data, lo, hi)
	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)
}

func partition(data sort.Interface, lo, hi int) (j int) {
	i, j := lo + 1, hi
	for {
		for i < hi && data.Less(i, lo) {
			i++
		}
		for j > lo && data.Less(lo, j) {
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
		i++
		j--
	}
	data.Swap(lo, j)
	return
}
