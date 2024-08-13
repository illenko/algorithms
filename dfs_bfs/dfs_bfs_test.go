package dfs_bfs

import (
	"testing"

	"reflect"
)

func TestDFS(t *testing.T) {
	tests := []struct {
		graph  map[string][]string
		start  string
		wanted string
		want   []string
	}{
		{
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {"D", "E"},
				"C": {"F"},
				"D": {},
				"E": {"F"},
				"F": {},
			},
			start:  "A",
			wanted: "F",
			want:   []string{"A", "C", "F"},
		},
		{
			graph: map[string][]string{
				"1": {"2", "3"},
				"2": {"4"},
				"3": {"4"},
				"4": {},
			},
			start:  "1",
			wanted: "4",
			want:   []string{"1", "3", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.start, func(t *testing.T) {
			got := dfs(tt.graph, tt.start, tt.wanted)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		graph  map[string][]string
		start  string
		wanted string
		want   []string
	}{
		{
			graph: map[string][]string{
				"A": {"B", "C"},
				"B": {"D", "E"},
				"C": {"F"},
				"D": {},
				"E": {"F"},
				"F": {},
			},
			start:  "A",
			wanted: "F",
			want:   []string{"A", "B", "C", "D", "E", "F"},
		},
		{
			graph: map[string][]string{
				"1": {"2", "3"},
				"2": {"4"},
				"3": {"4"},
				"4": {},
			},
			start:  "1",
			wanted: "4",
			want:   []string{"1", "2", "3", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.start, func(t *testing.T) {
			got := bfs(tt.graph, tt.start, tt.wanted)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
