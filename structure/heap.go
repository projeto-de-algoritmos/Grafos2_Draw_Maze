package structure

type PathHeap []Edge

func (h PathHeap) Len() int { return len(h)}

func (h PathHeap) Less(i, j int) bool {
	return h[i].Wgt < h[j].Wgt
}
func (h PathHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func (h *PathHeap) Push(x any) {
	edge := x.(Edge)
	*h = append(*h, edge)
}

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	edge := old[n-1]
	*h = old[0 : n-1]
	return edge
}
