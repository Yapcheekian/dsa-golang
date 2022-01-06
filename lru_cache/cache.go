package lrucache

type LRUCache struct {
	m map[int]int
	q []int
	c int
}

func Constructor(capacity int) *LRUCache {
	m := make(map[int]int)
	q := make([]int, 0, capacity)
	return &LRUCache{
		m: m,
		q: q,
		c: capacity,
	}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.m[key]; !ok {
		return -1
	}
	for i := 0; i < len(c.q); i++ {
		if c.q[i] == key {
			c.q = append(c.q[:i], c.q[i+1:]...)
			c.q = append([]int{key}, c.q...)
			break
		}
	}
	return c.m[key]

}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.m[key]; !ok {
		if len(c.q) == c.c {
			delete(c.m, c.q[len(c.q)-1])
			c.q = c.q[:len(c.q)-1]
		}
	} else {
		for i := 0; i < len(c.q); i++ {
			if c.q[i] == key {
				c.q = append(c.q[:i], c.q[i+1:]...)
				break
			}
		}
	}

	c.q = append([]int{key}, c.q...)
	c.m[key] = value
}
