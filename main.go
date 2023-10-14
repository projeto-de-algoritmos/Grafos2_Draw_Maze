package main

import "fmt"

func main() {
	graph :=Graph{}

	for i := 0; i < 8; i++ {
			graph.AddVertex(i)
	}

	graph.AddEdge(0, 2, 9)
	graph.AddEdge(0, 6, 14)
	graph.AddEdge(0, 7, 15)
	graph.AddEdge(2, 3, 23)
	graph.AddEdge(3, 1, 19)
	graph.AddEdge(3, 5, 2)
	graph.AddEdge(4, 3, 6)
	graph.AddEdge(4, 1, 6)
	graph.AddEdge(5, 4, 11)
	graph.AddEdge(5, 1, 16)
	graph.AddEdge(6, 3, 18)
	graph.AddEdge(6, 5, 30)
	graph.AddEdge(6, 7, 5)
	graph.AddEdge(7, 5, 20)
	graph.AddEdge(7, 1, 44)

	graph.Print()


	fmt.Println("Caminho:", Djikstra(graph, 0, 1))

}
