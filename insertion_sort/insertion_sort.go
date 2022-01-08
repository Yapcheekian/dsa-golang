package insertionsort

func InsertionSort(a []int) []int {
	l := len(a)
	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if less(a[j], a[j-1]) {
				swap(a, j, j-1)
			} else {
				break
			}
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
