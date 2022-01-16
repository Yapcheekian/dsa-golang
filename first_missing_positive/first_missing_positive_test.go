package firstmissingpositive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	input := []int{1, 2, 0}
	result := FirstMissingPositive(input)
	assert.Equal(t, 3, result)

	input = []int{2, 1, 3, 7, 6, 8, -1, -10, 15}
	result = FirstMissingPositive(input)
	assert.Equal(t, 4, result)

	input = []int{2, 3, -7, 6, 8, 1, -10, 15}
	result = FirstMissingPositive(input)
	assert.Equal(t, 4, result)

	input = []int{1, 1, 0, -1, -2}
	result = FirstMissingPositive(input)
	assert.Equal(t, 2, result)
}
