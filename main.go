package main

import (
)

func main() {
	graph :=Graph{}

	for i := 0; i < 6; i++ {
			graph.AddVertex(i)
	}
	graph.AddEdge(0, 1, 23)
	graph.AddEdge(1, 2, 23)
	graph.AddEdge(0, 2, 23)
	graph.AddEdge(3, 1, 23)
	graph.AddEdge(4, 5, 23)
	graph.AddEdge(5, 1, 23)

	graph.Print()

}
