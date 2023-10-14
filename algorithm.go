package main

import (
	"container/heap"
)

func incrementPath(ng *[]Edge, wgt int) {
	for i := range *ng {
		(*ng)[i].wgt += wgt	
	}
}

func Djikstra (g Graph, src, dst int) (int) {
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
			return path.wgt
		}

		neighbors = g.nodes[path.dst].ng
		incrementPath(&neighbors, path.wgt)
	}
	return -1
}
