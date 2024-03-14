package main

import "fmt"

type node struct {
	Value int
	Left  *node
	Right *node
}

func newNode(value int) *node {
	return &node{Value: value}
}

func main() {
	root := newNode(5)
	root.Left = newNode(3)
	root.Right = newNode(6)
	root.Left.Left = newNode(9)
	root.Right.Right = newNode(10)
	fmt.Println("In-order tranversal: ")
	inOrder(root)
}

func inOrder(node *node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Println(node.Value)
		inOrder(node.Right)
	}
}
