package addbinary

import "strconv"

func AddBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	var result []int
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = append([]int{sum % 2}, result...)
		carry = sum / 2
	}
	if carry != 0 {
		result = append([]int{carry}, result...)
	}

	var resultStr string
	for _, el := range result {
		str := strconv.Itoa(el)
		resultStr += str
	}
	return resultStr
}
