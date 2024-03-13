package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *stack) pop() int {
	var lastIndex = (len(s.items) - 1)
	var removedItem = s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return removedItem
}

func main() {
	var myStack = stack{}
	fmt.Println(myStack.items)
	myStack.push(1)
	myStack.push(150)
	myStack.push(50)
	fmt.Println(myStack.items)
	myStack.pop()
	fmt.Println(myStack.items)
	myStack.pop()
	fmt.Println(myStack.items)
}
