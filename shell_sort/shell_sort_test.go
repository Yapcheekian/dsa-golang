package shellsort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShellSort(t *testing.T) {
	input := []int{3, 4, 2, 7, 4, 1, 9, 7, 8, 6, 5, 1, 2, 3, 4}
	output := ShellSort(input)
	expOutput := []int{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 7, 7, 8, 9}
	for i := 0; i < len(output); i++ {
		assert.Equal(t, expOutput[i], output[i])
	}
}

func BenchmarkShellSort(b *testing.B) {
	input := generateArray()
	for i := 0; i < b.N; i++ {
		ShellSort(input)
	}
}

func generateArray() []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(100)
}
