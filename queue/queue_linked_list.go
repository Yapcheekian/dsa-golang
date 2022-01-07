package queue

type QueueLinkedList struct {
	first *node
	last  *node
}

type node struct {
	item string
	next *node
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{}
}

func (q *QueueLinkedList) Enqueue(item string) {
	if q.first == nil {
		n := &node{
			item: item,
		}
		q.first = n
		q.last = n
	} else {
		n := &node{
			item: item,
		}
		q.last.next = n
		q.last = n

	}
}

func (q *QueueLinkedList) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}
	n := q.first
	q.first = q.first.next
	return n.item
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.first == nil
}
