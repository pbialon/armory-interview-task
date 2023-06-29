package priority_queue

type Item interface {
	Lt(other interface{}) bool
}

type PriorityQueue struct {
	items []Item
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	first := pq.items[i]
	second := pq.items[j]
	return first.Lt(second)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.items = append(pq.items, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	if len(pq.items) == 0 {
		return nil
	}
	old := pq.items
	item := old[0]
	pq.items = old[1:]
	return item
}
