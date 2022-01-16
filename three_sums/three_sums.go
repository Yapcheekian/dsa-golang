package threesums

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			j := i + 1
			k := len(nums) - 1
			target := 0 - nums[i]
			for j < k {
				if nums[j]+nums[k] < target {
					j++
				} else if nums[j]+nums[k] > target {
					k--
				} else {
					res := []int{nums[i], nums[j], nums[k]}
					results = append(results, res)
					for j < k && nums[j] == nums[j+1] {
						j++
					}
					for j < k && nums[k] == nums[k-1] {
						k--
					}
					j++
					k--
				}
			}
		}
	}
	fmt.Println(results)
	return results
}
