package data_structures

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
