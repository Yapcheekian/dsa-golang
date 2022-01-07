package binarysearch

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if target > arr[mid] {
			low = mid + 1
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
