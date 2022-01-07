package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	assert.Equal(t, -1, c.Get(1))
	assert.Equal(t, 3, c.Get(2))
}
