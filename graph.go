package main

import "fmt"

type Edge struct {
	src, dst, wgt int
}

type Vertex struct {
	Key int
	ng []Edge
}

type Graph struct {
	size int
	nodes []Vertex
}

func (g *Graph) AddVertex(n int) {
	v  := Vertex{Key: n}
	g.nodes = append(g.nodes, v)
	g.size++
}

func (g *Graph) AddEdge(e Edge) {
	g.nodes[e.src].ng = append(g.nodes[e.src].ng, e)
}

func (g *Graph) Print() {
	for i, node := range g.nodes {
		fmt.Printf("%d\t| %d:", i, node.Key)
		for _, e := range node.ng {
			fmt.Printf("- %d -> %d\t|", e.wgt, g.nodes[e.dst].Key)
		}
		fmt.Print("\n")
	}
}
