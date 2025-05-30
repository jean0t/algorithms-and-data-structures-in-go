package data_structures

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(i int) {
	q.Items = append(q.Items, i)
}

func (q *Queue) Dequeue() int {
	var removedItem = q.Items[0]
	q.Items = q.Items[1:]
	return removedItem
}

func NewQueue() *Queue {
    return &Queue{Items: make([]int, 0)}
}
