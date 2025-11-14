//package binaryTree
package main

import "fmt"

type node struct {
	Value int
	Left *node
	Right *node
}

func newNode(data int) *node {
	return &node{Value: data}
}

func (n *node) insert(data int) {
	if data < n.Value {
		if n.Left == nil {
			n.Left =  newNode(data)
			return
		}
		n.Left.insert(data)
		
	} else if data > n.Value {
		if n.Right == nil {
			n.Right = newNode(data)
			return
		}
		n.Right.insert(data)
		
	} else {
		// nous n'avons pas accepter les dupliqués
		return
	}
}

type Tree struct {
	Node *node
}

func (t *Tree) Insert(data int) {
	if t.Node == nil {
		t.Node = newNode(data)
	}

	t.Node.insert(data)
}

func (t *Tree) Bfs() []int {
	var queue []*node
	var visited []int
	var current *node

	if t.Node == nil {
		return visited
	}

	queue = append(queue, t.Node)
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]

		visited = append(visited, current.Value)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return visited
}



func NewTree(value int) *Tree {
	return &Tree{Node: newNode(value)}
}

func main() {
	a := NewTree(2)
	a.Insert(1)
	a.Insert(3)
	a.Insert(5)
	a.Insert(4)
	fmt.Println(a.Node.Value)
	fmt.Println(a.Node.Left.Value)
	fmt.Println(a.Bfs())
}
