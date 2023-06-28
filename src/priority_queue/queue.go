package priority_queue

type PriorityQueue[Item ordered] struct {
	items []*Item
}

func (pq *PriorityQueue[Item]) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue[Item]) Less(i, j int) bool {
	return (*pq.items[i]).lt(*pq.items[j])
}

func (pq *PriorityQueue[Item]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue[Item]) Push(x interface{}) {
	pq.items = append(pq.items, x.(*Item))
}

func (pq *PriorityQueue[Item]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[Item]) Peek() interface{} {
	return pq.items[0]
}
