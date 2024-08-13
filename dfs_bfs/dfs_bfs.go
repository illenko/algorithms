package dfs_bfs

import "slices"

// DFS is a depth-first search algorithm.
// It uses a stack to keep track of the nodes to visit next.
// DFS is more efficient than BFS for visiting all nodes.
// O(V+E) time complexity where V is the number of vertices and E is the number of edges.
func dfs(graph map[string][]string, start string, wanted string) []string {
	var visited []string
	var stack []string
	stack = append(stack, start)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !slices.Contains(visited, node) {
			visited = append(visited, node)
			if node == wanted {
				break
			}
			stack = append(stack, graph[node]...)
		}
	}

	return visited
}

// BFS is a breadth-first search algorithm.
// It uses a queue to keep track of the nodes to visit next.
// BFS is more efficient than DFS for finding the shortest path.
// O(V+E) time complexity where V is the number of vertices and E is the number of edges.
func bfs(graph map[string][]string, start string, wanted string) []string {
	var visited []string
	var queue []string
	queue = append(queue, start)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !slices.Contains(visited, node) {
			visited = append(visited, node)
			if node == wanted {
				break
			}
			queue = append(queue, graph[node]...)
		}
	}

	return visited
}
