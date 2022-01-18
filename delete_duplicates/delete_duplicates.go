package deleteduplicates

type node struct {
	Val  int
	Next *node
}

func DeleteDuplicates(head *node) *node {
	if head == nil {
		return head
	}
	n := head
	for n != nil {
		next := n.Next
		for next != nil && n.Val == next.Val {
			next = next.Next
		}
		n.Next = next
		n = next
	}
	return head
}
