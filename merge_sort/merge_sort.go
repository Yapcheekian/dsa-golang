package mergesort

func MergeSort(a []int) {
	aux := make([]int, len(a))
	sort(a, aux, 0, len(a)-1)
}

func sort(a []int, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(a, aux, lo, mid)
	sort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a []int, aux []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}

func less(x, y int) bool {
	return x < y
}
