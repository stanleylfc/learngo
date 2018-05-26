package main

var graph = make(map[string]map[string]bool)

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func setEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}
