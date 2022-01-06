package main

import (
	"fmt"

	LRUCache "github.com/Yapcheekian/dsa-golang/lru_cache"
)

func main() {
	c := LRUCache.Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
