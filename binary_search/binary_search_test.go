package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	input := []int{1, 5, 8, 10, 15, 20, 35}
	target := 5
	expResult := 1
	result := BinarySearch(input, target)
	assert.Equal(t, expResult, result)

	target = 6
	expResult = -1
	result = BinarySearch(input, target)
	assert.Equal(t, expResult, result)

	target = 35
	expResult = 6
	result = BinarySearch(input, target)
	assert.Equal(t, expResult, result)
}
