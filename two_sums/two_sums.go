package twosums

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		remainder := target - num
		if _, ok := m[remainder]; ok {
			return []int{i, m[remainder]}
		}
		m[num] = i
	}
	return nil
}
