package merge_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		array    []int
		expected []int
	}{
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{5, 3, 8, 6, 2}, []int{2, 3, 5, 6, 8}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := mergeSort(test.array)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("quickSort(%v) = %v; want %v", test.array, result, test.expected)
		}
	}
}
