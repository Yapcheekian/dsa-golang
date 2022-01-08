package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryHeap(t *testing.T) {
	h := New()
	h.Insert(1)
	h.Insert(5)
	h.Insert(2)
	h.Insert(3)
	h.Insert(6)
	expResult := []int{0, 6, 5, 2, 1, 3}
	for i := 0; i < len(h.arr); i++ {
		assert.Equal(t, expResult[i], h.arr[i])
	}
	result := h.DeleteMax()
	assert.Equal(t, 6, result)
	result = h.DeleteMax()
	assert.Equal(t, 5, result)
	result = h.DeleteMax()
	assert.Equal(t, 3, result)
	result = h.DeleteMax()
	assert.Equal(t, 2, result)
}
