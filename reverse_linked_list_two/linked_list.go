package linkedlist

type node struct {
	Val  int
	Next *node
}

func ReverseBetween(head *node, left int, right int) *node {
	if left == right {
		return head
	}
	var stack []*node
	var start *node
	n := head

	if left != 1 {
		for i := 1; i < left-1; i++ {
			n = n.Next
		}
		start = n
		n = n.Next
	}

	for i := left; i <= right && n != nil; i++ {
		stack = append(stack, n)
		n = n.Next
	}

	l1 := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if left == 1 {
		head = l1
	} else {
		start.Next = l1
	}

	for len(stack) != 0 {
		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		l1.Next = next
		l1 = next
	}
	l1.Next = n

	return head
}
