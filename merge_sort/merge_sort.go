package merge_sort

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	left := array[:mid]
	right := array[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}
