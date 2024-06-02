package main

import (
	"fmt"
)


type node struct {
	next *node
	data int
}

type linkedList struct {
	head *node
	length int
}


func (l *linkedList) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
	  newNode.next = l.head
	}
	l.head = &newNode
	l.length++
  }

  func (l *linkedList) printList() {
	if l.head == nil {
	  return
	}
  
	currentNode := l.head
  
	for currentNode != nil {
	  fmt.Print(currentNode.data, " -> ")
	  currentNode = currentNode.next
	}
	fmt.Println("nil") 
  }


  func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next != nil && previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	if previousToDelete.next == nil {
		return
	  }
	previousToDelete.next = previousToDelete.next.next
	l.length--

  }

func main() {
	list := linkedList{}
	list.prepend(5)
	list.prepend(3)
	list.prepend(1)
	list.prepend(8)

	fmt.Println("Linked List")
	list.printList()
	list.deleteWithValue(3)
	list.printList()
}
