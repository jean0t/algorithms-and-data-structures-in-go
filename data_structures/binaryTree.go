package data_structures

type node struct {
	Value int
	Left  *node
	Right *node
}

func newNode(value int) *node {
	return &node{Value: value}
}
