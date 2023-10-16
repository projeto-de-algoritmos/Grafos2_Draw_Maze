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
	Nodes []Vertex
}

func (g *Graph) AddVertex(n int) {

	v  := Vertex{Key: n}

	if(g.Nodes[n].isEmpty()) {
		g.Nodes = append(g.Nodes, v)
		g.Size++
	}
}

func (e Vertex) isEmpty() bool {
	return (e.Key == 0)
}

func (g *Graph) AddDEdge(e Edge) {
	g.Nodes[e.Src].Ng = append(g.Nodes[e.Src].Ng, e)
}

func (g *Graph) AddEdge(e Edge) {
	ed := Edge{Src: e.Dst, Dst: e.Src, Wgt: e.Wgt}
	g.Nodes[e.Src].Ng = append(g.Nodes[e.Src].Ng, e)
	g.Nodes[e.Dst].Ng = append(g.Nodes[e.Dst].Ng, ed)
}

func (g *Graph) Init(size int) {
	g.Nodes = make([]Vertex, size, size)
}

func (g *Graph) Print() {
	for i:= 0; i< g.Size; i++ {
		fmt.Printf("%d\t| ", i)
		for _, e := range g.Nodes[i].Ng {
			fmt.Printf("%d -> %d\t|", e.Wgt, e.Dst)
		}
		fmt.Print("\n")
	}
}
