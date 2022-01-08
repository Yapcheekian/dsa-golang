package insertionsort

func InsertionSort(a []int) []int {
	l := len(a)
	for i := 1; i < l; i++ {
		for j := i; j > 0 && less(a[j], a[j-1]); j-- {
			swap(a, j, j-1)
		}
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
