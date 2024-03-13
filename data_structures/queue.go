package main

import "fmt"

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

func main() {
	var myQueue = queue{}
	fmt.Println(myQueue.items)
	myQueue.enqueue(50)
	myQueue.enqueue(37)
	myQueue.enqueue(23)
	myQueue.enqueue(11)
	fmt.Println(myQueue.items)
	myQueue.dequeue()
	myQueue.dequeue()
	fmt.Println(myQueue.items)
}
