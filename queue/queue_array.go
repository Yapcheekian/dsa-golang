package queue

type QueueArray struct {
	arr []string
}

func NewQueueArray() *QueueArray {
	return &QueueArray{
		arr: make([]string, 0, 5),
	}
}

func (q *QueueArray) Enqueue(item string) {
	q.arr = append(q.arr, item)
}

func (q *QueueArray) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}
	item := q.arr[0]
	q.arr = q.arr[1:]
	return item
}

func (q *QueueArray) IsEmpty() bool {
	return len(q.arr) == 0
}
