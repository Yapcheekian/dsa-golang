package binaryheap

import "fmt"

type BinaryHeap struct {
	arr []int
}

func New() *BinaryHeap {
	arr := make([]int, 1, 5)
	return &BinaryHeap{
		arr: arr,
	}
}

func (h *BinaryHeap) Insert(x int) {
	h.arr = append(h.arr, x)
	h.swim(len(h.arr) - 1)
}

func (h *BinaryHeap) DeleteMax() int {
	max := h.arr[1]
	h.swap(1, len(h.arr)-1)
	h.arr = h.arr[:len(h.arr)-1]
	h.sink(1)
	return max
}

func (h *BinaryHeap) swim(k int) {
	for k > 1 && less(h.arr[k/2], h.arr[k]) {
		h.swap(k/2, k)
		k /= 2
	}
}

func (h *BinaryHeap) sink(k int) {
	for 2*k+1 < len(h.arr) {
		j := 2 * k
		if j < len(h.arr) && less(h.arr[j], h.arr[j+1]) {
			j++
		}
		if !less(h.arr[k], h.arr[j]) {
			break
		}
		h.swap(k, j)
		k = j
	}
}

func less(x, y int) bool {
	return x < y
}

func (h *BinaryHeap) swap(x, y int) {
	tmp := h.arr[x]
	h.arr[x] = h.arr[y]
	h.arr[y] = tmp
}

func debug(h *BinaryHeap) {
	fmt.Println(h.arr)
}
