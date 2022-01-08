package shellsort

func ShellSort(a []int) []int {
	l := len(a)
	h := 1
	for h < l/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h && less(a[j], a[j-h]); j -= h {
				swap(a, j, j-h)
			}
		}
		h = h / 3
	}
	return a
}

func less(x, y int) bool {
	return x < y
}

func swap(a []int, x, y int) {
	tmp := a[x]
	a[x] = a[y]
	a[y] = tmp
}
