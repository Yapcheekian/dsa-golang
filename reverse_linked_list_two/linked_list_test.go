package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	n1 := &node{Val: 1}
	n2 := &node{Val: 2}
	n3 := &node{Val: 3}
	n4 := &node{Val: 4}
	n5 := &node{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	result := ReverseBetween(n1, 2, 4)
	expResult := []int{1, 4, 3, 2, 5}
	i := 0
	for result != nil {
		assert.Equal(t, expResult[i], result.Val)
		i++
		result = result.Next
	}

	n1 = &node{Val: 5}
	result = ReverseBetween(n1, 1, 1)
	expResult = []int{5}
	i = 0
	for result != nil {
		assert.Equal(t, expResult[i], result.Val)
		i++
		result = result.Next
	}

	n1 = &node{Val: 3}
	n2 = &node{Val: 5}
	n1.Next = n2
	result = ReverseBetween(n1, 1, 2)
	expResult = []int{5, 3}
	i = 0
	for result != nil {
		assert.Equal(t, expResult[i], result.Val)
		i++
		result = result.Next
	}
}
