package binaryTree

type node struct {
	Value int
	Left  *node
	Right *node
}

func newNode(value int) *node {
	return &node{Value: value}
}

func (n *node) insert(data int) {
	if data < n.Value {
		if n.Left == nil {
			n.Left = newNode(data)
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

func (n *node) search(data int) bool {
	if data == n.Value {
		return true
	}

	if data < n.Value {
		if n.Left == nil {
			return false
		}
		n.Left.search(data)

	} else if data > n.Value {
		if n.Right == nil {
			return false
		}

		n.Right.search(data)
	}
}

type Tree struct {
	Node *node
}

func (t *Tree) Insert(data int) {
	if t.Node == nil {
		t.Node = &node{Value: data}
	}

	t.Node.insert(data)
}

func (t *Tree) Search(data int) bool {
	if t.Node == nil {
		return nil
	}

	return t.Node.search(data)
}
