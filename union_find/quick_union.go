package unionfind

type QuickUnion struct {
	arr []int
}

func NewQuickUnion(n int) *QuickUnion {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return &QuickUnion{
		arr: arr,
	}
}

func (q *QuickUnion) Connected(x, y int) bool {
	return q.root(x) == q.root(y)
}

func (q *QuickUnion) Union(x, y int) {
	rootX := q.root(x)
	rootY := q.root(y)
	q.arr[rootX] = rootY
}

func (q *QuickUnion) root(i int) int {
	for i != q.arr[i] {
		i = q.arr[i]
	}
	return i
}
