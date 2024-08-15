package main

// DijkstraSearch is a function that implements the Dijkstra search algorithm.
// It finds the shortest path between two nodes in a graph.
// O(V^2) time complexity where V is the number of vertices.
// O(V) space complexity where V is the number of vertices.
// It returns the shortest path between the start and end nodes.
// The graph is represented as a map of nodes to a map of neighbors and their costs.
func dijkstraSearch(graph map[string]map[string]int, start, end string) []string {
	costs := make(map[string]int)
	parents := make(map[string]string)
	processed := make(map[string]bool)

	for node, cost := range graph[start] {
		costs[node] = cost
		parents[node] = start
	}

	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n, c := range neighbors {
			newCost := cost + c
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs, processed)
	}

	path := []string{end}
	for path[len(path)-1] != start {
		path = append(path, parents[path[len(path)-1]])
	}

	return reverse(path)
}

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := int(^uint(0) >> 1)
	lowestCostNode := ""
	for node, cost := range costs {
		if cost < lowestCost && !processed[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func reverse(array []string) []string {
	for i := 0; i < len(array)/2; i++ {
		j := len(array) - i - 1
		array[i], array[j] = array[j], array[i]
	}
	return array
}
