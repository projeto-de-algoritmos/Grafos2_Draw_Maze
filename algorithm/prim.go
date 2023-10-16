package algorithm

import (
	"container/heap"
	"fmt"
	"proj/structure"
)

func Prim (g structure.Graph) {
	pq := structure.PathHeap{}
	heap.Init(&pq)

	visited := make([]bool, g.Size)
	dist := make([]structure.Edge, 0, g.Size)
	var wgt int

	visited[0] = true
	ng := g.Nodes[0].Ng
	s := g.Size-1
	for s!=0{
		for _, e := range ng {
			if visited[e.Dst] == false {
				heap.Push(&pq, e)
			}
		}
		path := heap.Pop(&pq).(structure.Edge)
		wgt += path.Wgt
		visited[path.Dst] = true
		dist = append(dist, path)
		s--
		if !(pq.Len() > 0) {
			break
		}
		ng = g.Nodes[path.Dst].Ng
	}
	fmt.Println("A MST é:")
	fmt.Print(dist)
	fmt.Println("Peso:", wgt)
}
