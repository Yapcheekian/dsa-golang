package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStackLinkedList()
	assert.Equal(t, true, s.IsEmpty())
	s.Push("a")
	s.Push("b")
	s.Push("c")
	result := s.Pop()
	assert.Equal(t, "c", result)
	assert.Equal(t, false, s.IsEmpty())

	sa := NewStackLinkedList()
	assert.Equal(t, true, sa.IsEmpty())
	sa.Push("a")
	sa.Push("b")
	sa.Push("c")
	result = sa.Pop()
	assert.Equal(t, "c", result)
	assert.Equal(t, false, sa.IsEmpty())
}
