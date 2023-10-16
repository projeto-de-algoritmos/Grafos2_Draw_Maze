package algorithm

import (
	"container/heap"
	"fmt"
	"proj/structure"
)

func incrementPath(ng *[]structure.Edge, wgt int) {
	for i := range *ng {
		(*ng)[i].Wgt += wgt	
	}
}

func printPath(e []structure.Edge, src, dst int) {
	if src != dst {
		printPath(e, src, e[dst].Src)
		fmt.Printf(" -> %d", e[dst].Dst)
	} else {
		fmt.Printf("Caminho: %d", src)
	}

}

func Djikstra (g structure.Graph, src, dst int) {
	pq := structure.PathHeap{}
	heap.Init(&pq)	

	visited := make([]bool, g.Size)
	visited[src] = true

	dist := make([]structure.Edge, g.Size)	

	neighbors := g.Nodes[src].Ng

	for {
		for _, e := range neighbors {
			if visited[e.Dst] == false {
				heap.Push(&pq, e)
			}
		}

		path := heap.Pop(&pq).(structure.Edge)
		visited[path.Dst] = true
		dist[path.Dst] = path

		if !(pq.Len() > 0) {
			break
		}

		if path.Dst == dst {
			println("Caminho de ",src,"até ",dst)
			printPath(dist, src, dst)
			fmt.Printf("\n")
			fmt.Println("Tamanho: ", path.Wgt)
			return
		}

		neighbors = g.Nodes[path.Dst].Ng
		incrementPath(&neighbors, path.Wgt)
	}
	fmt.Println("Não existe caminho :/")
}
