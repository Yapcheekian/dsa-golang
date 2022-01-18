package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	input := []int{2, 2, 1}
	result := SingleNumber(input)
	assert.Equal(t, 1, result)

	input = []int{4, 1, 2, 1, 2}
	result = SingleNumber(input)
	assert.Equal(t, 4, result)

	input = []int{1}
	result = SingleNumber(input)
	assert.Equal(t, 1, result)
}
