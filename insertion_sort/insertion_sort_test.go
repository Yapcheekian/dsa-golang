package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	input := []int{3, 4, 2, 7, 4, 1, 9, 7, 8}
	output := InsertionSort(input)
	expOutput := []int{1, 2, 3, 4, 4, 7, 7, 8, 9}
	for i := 0; i < len(output); i++ {
		assert.Equal(t, expOutput[i], output[i])
	}
}
