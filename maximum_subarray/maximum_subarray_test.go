package maximumsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := MaxSubArray(input)
	assert.Equal(t, 6, result)

	// input = []int{1}
	// result = MaxSubArray(input)
	// assert.Equal(t, 1, result)

	// input = []int{5, 4, -1, 7, 8}
	// result = MaxSubArray(input)
	// assert.Equal(t, 23, result)
}
