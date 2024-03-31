package main

import "fmt"

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

func main(){
  myList := linkedList{}
  node1 := &node{data:48}
  node2 := &node{data:18}
  node3 := &node{data:16}
  node4 := &node{data:67}
  node5 := &node{data:31}
  myList.append(node1)
  myList.append(node2)
  myList.append(node3)
  myList.append(node4)
  myList.append(node5)

  myList.print()
  myList.delete(16)
  myList.print()
  myList.delete(100)
  myList.print()
  myList.delete(48)
  emptyList := linkedList{}
  emptyList.delete(10)
  myList.print()
}
