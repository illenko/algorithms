package binary_search

func binarySearch(array []int, item int) int {

	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] == item {
			return mid
		} else if array[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
