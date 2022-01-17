package maximumsubarray

func MaxSubArray(nums []int) int {
	curr := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		curr += nums[i]
		if curr < 0 || nums[i] > curr {
			curr = nums[i]
		}
		if res < curr {
			res = curr
		}
	}
	return res
}
