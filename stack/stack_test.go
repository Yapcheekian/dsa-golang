package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	sLinkedList := NewStackLinkedList()
	assert.Equal(t, true, sLinkedList.IsEmpty())
	assert.Equal(t, "", sLinkedList.Pop())
	sLinkedList.Push("a")
	sLinkedList.Push("b")
	sLinkedList.Push("c")
	result := sLinkedList.Pop()
	assert.Equal(t, "c", result)
	assert.Equal(t, false, sLinkedList.IsEmpty())

	sArray := NewStackLinkedList()
	assert.Equal(t, true, sArray.IsEmpty())
	assert.Equal(t, "", sArray.Pop())
	sArray.Push("a")
	sArray.Push("b")
	sArray.Push("c")
	result = sArray.Pop()
	assert.Equal(t, "c", result)
	assert.Equal(t, false, sArray.IsEmpty())
}
