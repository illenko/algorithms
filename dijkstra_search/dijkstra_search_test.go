package main

import (
	"reflect"
	"testing"
)

func TestDijkstraSearch(t *testing.T) {
	tests := []struct {
		graph map[string]map[string]int
		start string
		end   string
		want  []string
	}{
		{
			graph: map[string]map[string]int{
				"A": {"B": 1, "C": 4},
				"B": {"C": 2, "D": 5},
				"C": {"D": 1},
				"D": {},
			},
			start: "A",
			end:   "D",
			want:  []string{"A", "B", "C", "D"},
		},
		{
			graph: map[string]map[string]int{
				"1": {"2": 2, "3": 4},
				"2": {"3": 1, "4": 7},
				"3": {"4": 3},
				"4": {},
			},
			start: "1",
			end:   "4",
			want:  []string{"1", "2", "3", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.start+" to "+tt.end, func(t *testing.T) {
			got := dijkstraSearch(tt.graph, tt.start, tt.end)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dijkstraSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
