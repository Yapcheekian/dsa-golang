package dynamicprogramming

func Max(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	maxRemainder := Max(a[1:])

	if a[0] > maxRemainder {
		return a[0]
	} else {
		return maxRemainder
	}
}

func Fib(n int, m map[int]int) int {
	if n <= 1 {
		return n
	}

	if _, ok := m[n]; !ok {
		m[n] = Fib(n-1, m) + Fib(n-2, m)
	}

	return m[n]
}
