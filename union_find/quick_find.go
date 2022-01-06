package unionfind

type QuickFind struct {
	arr []int
}

func NewQuickFind(n int) *QuickFind {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return &QuickFind{
		arr: arr,
	}
}

func (q *QuickFind) Connected(x, y int) bool {
	return q.arr[x] == q.arr[y]
}

func (q *QuickFind) Union(x, y int) {
	idX := q.arr[x]
	idY := q.arr[y]

	for i := 0; i < len(q.arr); i++ {
		if q.arr[i] == idX {
			q.arr[i] = idY
		}
	}
}
