package deleteduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	n1 := &node{Val: 1}
	n2 := &node{Val: 1}
	n3 := &node{Val: 2}
	n1.Next = n2
	n2.Next = n3

	DeleteDuplicates(n1)
	expResult := []int{1, 2}
	var i int
	for n1 != nil {
		assert.Equal(t, expResult[i], n1.Val)
		i++
		n1 = n1.Next
	}

	n1 = &node{Val: 1}
	n2 = &node{Val: 1}
	n3 = &node{Val: 2}
	n4 := &node{Val: 3}
	n5 := &node{Val: 3}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	DeleteDuplicates(n1)
	expResult = []int{1, 2, 3}
	i = 0
	for n1 != nil {
		assert.Equal(t, expResult[i], n1.Val)
		i++
		n1 = n1.Next
	}

	n1 = &node{Val: 1}
	n2 = &node{Val: 1}
	n3 = &node{Val: 1}
	n1.Next = n2
	n2.Next = n3

	DeleteDuplicates(n1)
	expResult = []int{1}
	i = 0
	for n1 != nil {
		if i > len(expResult)-1 {
			t.Fatal("wrong answer")
		}
		assert.Equal(t, expResult[i], n1.Val)
		i++
		n1 = n1.Next
	}
}
