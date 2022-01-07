package twosums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSums(t *testing.T) {
	nums := []int{2, 7, 11, 5}
	target := 9
	result := TwoSum(nums, target)
	assert.ElementsMatch(t, []int{1, 0}, result)
}
