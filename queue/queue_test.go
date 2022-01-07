package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	qLinkedList := NewQueueLinkedList()
	assert.Equal(t, true, qLinkedList.IsEmpty())
	assert.Equal(t, "", qLinkedList.Dequeue())
	qLinkedList.Enqueue("a")
	qLinkedList.Enqueue("b")
	qLinkedList.Enqueue("c")
	assert.Equal(t, "a", qLinkedList.Dequeue())
	assert.Equal(t, "b", qLinkedList.Dequeue())
	assert.Equal(t, false, qLinkedList.IsEmpty())
	assert.Equal(t, "c", qLinkedList.Dequeue())
	assert.Equal(t, true, qLinkedList.IsEmpty())

	qArray := NewQueueArray()
	assert.Equal(t, true, qArray.IsEmpty())
	assert.Equal(t, "", qArray.Dequeue())
	qArray.Enqueue("a")
	qArray.Enqueue("b")
	qArray.Enqueue("c")
	assert.Equal(t, "a", qArray.Dequeue())
	assert.Equal(t, "b", qArray.Dequeue())
	assert.Equal(t, false, qArray.IsEmpty())
	assert.Equal(t, "c", qArray.Dequeue())
	assert.Equal(t, true, qArray.IsEmpty())
}
