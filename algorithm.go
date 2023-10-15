package main

import (
	"container/heap"
	"fmt"
)

func incrementPath(ng *[]Edge, wgt int) {
	for i := range *ng {
		(*ng)[i].wgt += wgt	
	}
}

func printPath(e []Edge, src, dst int) {
	if src != dst {
		printPath(e, src, e[dst].src)
		fmt.Printf(" -> %d", e[dst].src)
	} else {
		fmt.Printf("Caminho:")
	}

}

func Djikstra (g Graph, src, dst int) {
	pq := PathHeap{}
	heap.Init(&pq)	

	visited := make([]bool, g.size)
	visited[src] = true

	dist := make([]Edge, g.size)	

	neighbors := g.nodes[src].ng

	for {
		for _, e := range neighbors {
			if visited[e.dst] == false {
				heap.Push(&pq, e)
			}
		}

		path := heap.Pop(&pq).(Edge)
		visited[path.dst] = true
		dist[path.dst] = path

		if !(pq.Len() > 0) {
			break
		}

		if path.dst == dst {
			printPath(dist, src, dst)
			fmt.Printf("\n")
			fmt.Println("Tamanho: ", path.wgt)
			return
		}

		neighbors = g.nodes[path.dst].ng
		incrementPath(&neighbors, path.wgt)
	}
	fmt.Println("Não existe caminho :/")
}
