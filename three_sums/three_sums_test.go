package threesums

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	_ = ThreeSum(nums)
	_ = [][]int{{-1, -1, 2}, {-1, 0, 1}}
}
