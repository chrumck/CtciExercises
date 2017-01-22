package algorithms

// MergeSort sorts given data using MergeSort algorithm.
func MergeSort(data SortInterface) {
	n := data.Len()
	aux := make([]interface{}, n, n)

	copyValuesFromAux := func(lo, hi int) {
		for i := lo; i <= hi; i++ {
			data.SetValue(i, aux[i])
		}
	}

	mergeToAux := func(lo, mid, hi int) {
		i, j := lo, mid+1
		for k := lo; k <= hi; k++ {
			switch {
			case i > mid:
				aux[k] = data.GetValue(j)
				j++
			case j > hi:
				aux[k] = data.GetValue(i)
				i++
			case data.Less(j, i):
				aux[k] = data.GetValue(j)
				j++
			default:
				aux[k] = data.GetValue(i)
				i++
			}
		}
	}

	var sort func(lo, hi int)
	sort = func(lo, hi int) {
		mid := (lo + hi) / 2
		if lo < mid {
			sort(lo, mid)
		}
		if mid+1 < hi {
			sort(mid+1, hi)
		}
		mergeToAux(lo, mid, hi)
		copyValuesFromAux(lo, hi)
	}

	sort(0, n-1)

}
