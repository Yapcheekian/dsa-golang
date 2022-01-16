package firstmissingpositive

import (
	"math"
)

func FirstMissingPositive(nums []int) int {
	k := segregate(nums)
	for i := 0; i < k; i++ {
		val := int(math.Abs(float64(nums[i])))
		if val > 0 && val <= k {
			if nums[val-1] > 0 {
				nums[val-1] = nums[val-1] * -1
			}
		}
	}

	for i := 0; i < k; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return k + 1
}

func segregate(a []int) int {
	k := 0
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			a[k] = a[i]
			k++
		}
	}
	return k
}
