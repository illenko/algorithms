package selection_sort

func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
	return array
}
