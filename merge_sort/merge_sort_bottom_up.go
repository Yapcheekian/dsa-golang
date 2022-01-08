package mergesort

func MergeSortBottomUp(a []int) {
	l := len(a)
	aux := make([]int, l)
	for sz := 1; sz < l; sz += sz {
		for lo := 0; lo < l-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, l-1))
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
