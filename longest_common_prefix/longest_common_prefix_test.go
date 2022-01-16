package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	same := LongestCommonPrefix(strs)
	assert.Equal(t, "fl", same)
}
