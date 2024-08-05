package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		array    []int
		item     int
		expected int
	}{
		{[]int{-1, 0, 1, 2, 4}, 4, 4},
		{[]int{0, 2, 3, 4, 5}, 0, 0},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{-10, -5, 0, 3, 7}, 3, 3},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 0},
	}

	for _, test := range tests {
		result := binarySearch(test.array, test.item)
		if result != test.expected {
			t.Errorf("binarySearch(%v, %d) = %d; want %d", test.array, test.item, result, test.expected)
		}
	}
}
