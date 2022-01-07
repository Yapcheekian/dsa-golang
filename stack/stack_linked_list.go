package stack

import "fmt"

type StackLinkedList struct {
	node *node
}

type node struct {
	item string
	next *node
}

func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{}
}

func (s *StackLinkedList) Push(item string) {
	if s.node == nil {
		s.node = &node{
			item: item,
		}
	} else {
		n := &node{
			item: item,
		}
		n.next = s.node
		s.node = n
	}
}

func (s *StackLinkedList) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	n := s.node
	s.node = s.node.next
	return n.item
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.node == nil
}

func debug(s *StackLinkedList) {
	for s.node != nil {
		fmt.Println(s.node.item)
		s.node = s.node.next
	}
}
