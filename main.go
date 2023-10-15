package main

func main() {
	graph :=Graph{}

	for i := 0; i < 8; i++ {
			graph.AddVertex(i)
	}

	graph.AddEdge(Edge{0, 2, 9})
	graph.AddEdge(Edge{0, 6, 14})
	graph.AddEdge(Edge{0, 7, 15})
	graph.AddEdge(Edge{2, 3, 23})
	graph.AddEdge(Edge{3, 1, 19})
	graph.AddEdge(Edge{3, 5, 2})
	graph.AddEdge(Edge{4, 3, 6})
	graph.AddEdge(Edge{4, 1, 6})
	graph.AddEdge(Edge{5, 4, 11})
	graph.AddEdge(Edge{5, 1, 16})
	graph.AddEdge(Edge{6, 3, 18})
	graph.AddEdge(Edge{6, 5, 30})
	graph.AddEdge(Edge{6, 7, 5})
	graph.AddEdge(Edge{7, 5, 20})
	graph.AddEdge(Edge{7, 1, 44})

	graph.Print()

	Djikstra(graph, 0, 1)

}
