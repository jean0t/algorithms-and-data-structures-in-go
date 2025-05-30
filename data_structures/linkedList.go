package data_structures

type node struct {
  data int
  next *node
}

type linkedList struct{
  head *node
  length int
}

func (l *linkedList) append (n *node) {
  second := l.head
  l.head = n
  l.head.next = second
  l.length++
}

func (l linkedList) print() {
  toPrint := l.head
  for l.length != 0 {
    fmt.Printf("%d ", toPrint.data)
    toPrint = toPrint.next
    l.length--
  }
  fmt.Println("")
}

func (l *linkedList) delete(value int) {
  if l.length == 0 {
    return
  }

  head := l.head

  if head.data == value {
    l.head = head.next
    l.length--
    return
  }

  for head.next.data != value {
    if head.next.next == nil {
      return
    }
    head = head.next
  }
  head.next = head.next.next
  l.length--
}
