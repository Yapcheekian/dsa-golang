package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	input := []int{3, 4, 2, 7, 4, 1, 9, 7, 8, 6, 5, 1}
	MergeSort(input)
	expOutput := []int{1, 1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9}
	for i := 0; i < len(input); i++ {
		assert.Equal(t, expOutput[i], input[i])
	}
}
