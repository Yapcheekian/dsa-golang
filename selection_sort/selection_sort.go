package selectionsort

func SelectionSort(a []int) []int {
	l := len(a)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if less(a[j], a[min]) {
				min = j
			}
		}
		swap(a, i, min)
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
