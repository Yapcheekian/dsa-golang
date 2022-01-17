package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	input := []int{1, 3, 5, 6}
	result := SearchInsert(input, 5)
	assert.Equal(t, 2, result)

	input = []int{1, 3, 5, 6}
	result = SearchInsert(input, 2)
	assert.Equal(t, 1, result)

	input = []int{1, 3, 5, 6}
	result = SearchInsert(input, 7)
	assert.Equal(t, 4, result)
}
