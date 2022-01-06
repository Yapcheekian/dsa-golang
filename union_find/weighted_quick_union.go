package unionfind

type WeightedQuickUnion struct {
	arr  []int
	sarr []int
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	arr := make([]int, n)
	sarr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		sarr[i] = 1
	}
	return &WeightedQuickUnion{
		arr:  arr,
		sarr: sarr,
	}
}

func (q *WeightedQuickUnion) Connected(x, y int) bool {
	return q.root(x) == q.root(y)
}

func (q *WeightedQuickUnion) Union(x, y int) {
	rootX := q.root(x)
	rootY := q.root(y)
	if q.sarr[x] > q.sarr[y] {
		q.arr[rootY] = rootX
		q.sarr[x] += q.sarr[y]
	} else {
		q.arr[rootX] = rootY
		q.sarr[y] += q.sarr[x]
	}
}

func (q *WeightedQuickUnion) root(i int) int {
	for i != q.arr[i] {
		// path compression
		q.arr[i] = q.arr[q.arr[i]]
		i = q.arr[i]
	}
	return i
}
