package quick_sort

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	var less []int
	var greater []int

	for _, v := range array[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
