package quicksort

import (
	"math/rand"
	"time"
)

func Quicksort(a []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	sort(a, 0, len(a)-1)
}

func sort(a []int, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, lo, hi)
	sort(a, lo, j-1)
	sort(a, j+1, hi)
}

func partition(a []int, lo, hi int) int {
	i := lo + 1
	j := hi
	for {
		for less(a[i], a[lo]) {
			if i == hi {
				break
			}
			i++
		}
		for less(a[lo], a[j]) {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		swap(a, i, j)
	}
	swap(a, j, lo)
	return j
}

func less(x, y int) bool {
	return x < y
}

func swap(a []int, x, y int) {
	tmp := a[x]
	a[x] = a[y]
	a[y] = tmp
}
