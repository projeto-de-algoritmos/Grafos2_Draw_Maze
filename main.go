package main

import (
	"proj/algorithm"
	"proj/structure"
)

func main() {
	graph := structure.Graph{}
	edge := structure.Edge{Src:0, Dst: 6, Wgt: 9}

	graph.Init()
	for i := 0; i < 8; i++ {
			graph.AddVertex(i)
	}

	graph.AddEdge(edge)
	graph.AddEdge(structure.Edge{Src:0, Dst:6, Wgt:14})
	graph.AddEdge(structure.Edge{Src:0, Dst:7, Wgt:15})
	graph.AddEdge(structure.Edge{Src:2, Dst:3, Wgt:23})
	graph.AddEdge(structure.Edge{Src:3, Dst:1, Wgt:19})
	graph.AddEdge(structure.Edge{Src:3, Dst:5, Wgt:2})
	graph.AddEdge(structure.Edge{Src:4, Dst:3, Wgt:6})
	graph.AddEdge(structure.Edge{Src:4, Dst:1, Wgt:6})
	graph.AddEdge(structure.Edge{Src:5, Dst:4, Wgt:11})
	graph.AddEdge(structure.Edge{Src:5, Dst:1, Wgt:16})
	graph.AddEdge(structure.Edge{Src:6, Dst:3, Wgt:18})
	graph.AddEdge(structure.Edge{Src:6, Dst:5, Wgt:30})
	graph.AddEdge(structure.Edge{Src:6, Dst:7, Wgt:5})
	graph.AddEdge(structure.Edge{Src:7, Dst:5, Wgt:20})
	graph.AddEdge(structure.Edge{Src:7, Dst:1, Wgt:44})

	graph.Print()

	algorithm.Djikstra(graph, 0, 1)

	g :=structure.Graph{}
	g.Init()
	for i := 0; i < 6; i++ {
			g.AddVertex(i)
	}
	

	g.AddEdge(structure.Edge{Src:0, Dst:1, Wgt:2})
	g.AddEdge(structure.Edge{Src:0, Dst:2, Wgt:3})
	g.AddEdge(structure.Edge{Src:1, Dst:2, Wgt:5})
	g.AddEdge(structure.Edge{Src:1, Dst:3, Wgt:4})
	g.AddEdge(structure.Edge{Src:1, Dst:4, Wgt:4})
	g.AddEdge(structure.Edge{Src:2, Dst:4, Wgt:5})
	g.AddEdge(structure.Edge{Src:3, Dst:4, Wgt:2})
	g.AddEdge(structure.Edge{Src:3, Dst:5, Wgt:3})
	g.AddEdge(structure.Edge{Src:4, Dst:5, Wgt:5})

	g.AddEdge(structure.Edge{Src:1, Dst:0, Wgt:2})
	g.AddEdge(structure.Edge{Src:2, Dst:0, Wgt:3})
	g.AddEdge(structure.Edge{Src:2, Dst:1, Wgt:5})
	g.AddEdge(structure.Edge{Src:3, Dst:1, Wgt:4})
	g.AddEdge(structure.Edge{Src:4, Dst:1, Wgt:4})
	g.AddEdge(structure.Edge{Src:4, Dst:2, Wgt:5})
	g.AddEdge(structure.Edge{Src:4, Dst:3, Wgt:2})
	g.AddEdge(structure.Edge{Src:5, Dst:3, Wgt:3})
	g.AddEdge(structure.Edge{Src:5, Dst:4, Wgt:5})
	algorithm.Prim(g)

}
