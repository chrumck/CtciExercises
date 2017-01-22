package algorithms

import "sort"

func isSorted(data sort.Interface) bool {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

var sortTestSamples = []IntSlice{
   IntSlice{sort.IntSlice{3, 3, 3, 3, 3, 3, 3, 3}},
   IntSlice{sort.IntSlice{1, 2, 3, 4, 5, 6, 7}},
   IntSlice{sort.IntSlice{7, 6, 4, 3, 2, 1, 0, -1}},
   IntSlice{sort.IntSlice{33, 12, 21, 3, 6, 5, 100, 71, 7, 6, 4, 8, 2, 1, -100, -1}},
   IntSlice{sort.IntSlice{33, 12, 21, 21, 21, 5, 100, 71, 7, 6, 4, 8, 8, 8, -100, -1}},
   IntSlice{sort.IntSlice{8, 7}},
   IntSlice{sort.IntSlice{3}},
   IntSlice{sort.IntSlice{3, 2, 1}},
   IntSlice{sort.IntSlice{1, 2, 3}},
   IntSlice{sort.IntSlice{4, 1, 2, 3, 7}},
}
