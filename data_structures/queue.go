package data_structures

type queue struct {
	items []int
}

func (q *queue) enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *queue) dequeue() int {
	var removedItem = q.items[0]
	q.items = q.items[1:]
	return removedItem
}
