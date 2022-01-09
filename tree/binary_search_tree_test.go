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
}
