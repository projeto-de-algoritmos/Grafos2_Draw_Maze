package main

import (
	"fmt"
	"proj/algorithm"
	"proj/structure"
)

func main() {
	graph := structure.Graph{}

	fmt.Println("Insira 1: para Dijkstra, 2: para Prim")
	var opt int
	fmt.Scanf("%d", &opt)

	fmt.Println("Informe o número de nós do grafo.")
	var size int
	fmt.Scanf("%d", &size)
	graph.Init(size)

	for i := 0; i < size; i++ {
			graph.AddVertex(i)
	}

	fmt.Println("Informe os caminhos de A para B e seu peso. (Ex: 0 3 5); Se um dado for negativo a sessão é terminada.")

	for {
		var src, dst, wgt int
		fmt.Scanf("%d %d %d", &src, &dst, &wgt)
		if (src<0 || dst<0 || wgt<0){
			break
		} 
		e := structure.Edge{Src: src, Dst: dst, Wgt: wgt}
		if opt == 1 {
			graph.AddDEdge(e)
		} else {
			graph.AddEdge(e)
		}
	}


	if opt == 1 {
		fmt.Println("Informe o caminho (ex: 0 1)")
		var src, dst int
		graph.Print()
		fmt.Scanf("%d %d", &src, &dst)
		algorithm.Djikstra(graph, src, dst)
	} else {
		graph.Print()
		algorithm.Prim(graph)
	}

}
