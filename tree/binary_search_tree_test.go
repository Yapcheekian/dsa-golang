package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	bst := NewBST()
	bst.Set("name", "yap")
	assert.Equal(t, "yap", bst.Get("name"))
	assert.Equal(t, "", bst.Get("namexxx"))
	bst.Set("name", "yap2")
	assert.Equal(t, "yap2", bst.Get("name"))
	bst.Set("aa", "min")
	assert.Equal(t, "min", bst.Get("aa"))
	bst.DeleteMin()
	assert.Equal(t, "", bst.Get("aa"))
	bst.Set("zzz", "max")
	assert.Equal(t, "max", bst.Get("zzz"))
	bst.DeleteMax()
	assert.Equal(t, "", bst.Get("zzz"))
}
