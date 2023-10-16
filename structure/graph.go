package structure

import "fmt"

type Edge struct {
	Src, Dst, Wgt int
}

type Vertex struct {
	Key int
	Ng []Edge
}

type Graph struct {
	Size int
	//nodes []Vertex
	Nodes map[int]*Vertex
}

func (g *Graph) AddVertex(n int) {
	v  := Vertex{Key: n}
	//g.nodes = append(g.nodes, v)
	g.Nodes[n] = &v
	g.Size++
}

func (g *Graph) AddEdge(e Edge) {
	//g.nodes[e.src].ng = append(g.nodes[e.src].ng, e)
	g.Nodes[e.Src].Ng = append(g.Nodes[e.Src].Ng, e)
}

func (g *Graph) Init() {
	g.Nodes = make(map[int]*Vertex)
}
func (g *Graph) InitGraph(edges []Edge) {

	for i := 0; i<len(edges); i++ {
		g.AddVertex(i)	
	}
	for _, e := range edges {
		g.AddEdge(e)
	}
}

func (g *Graph) Print() {
	for i, node := range g.Nodes {
		fmt.Printf("%d\t| %d:", i, node.Key)
		for _, e := range node.Ng {
			fmt.Printf("- %d -> %d\t|", e.Wgt, g.Nodes[e.Dst].Key)
		}
		fmt.Print("\n")
	}
}
