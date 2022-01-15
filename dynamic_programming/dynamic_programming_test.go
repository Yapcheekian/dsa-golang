package dynamicprogramming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	result := Max([]int{3, 5, 8, 9, 1, 4, 7})
	assert.Equal(t, 9, result)
}

func TestFib(t *testing.T) {
	result := Fib(20, make(map[int]int))
	assert.Equal(t, 6765, result)
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20, make(map[int]int))
	}
}

func BenchmarkSlowFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowFib(20)
	}
}

func slowFib(n int) int {
	if n <= 1 {
		return n
	}
	return slowFib(n-1) + slowFib(n-2)
}
