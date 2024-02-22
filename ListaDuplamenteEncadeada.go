package main

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) append(data int) {
	newNode := &Node{data: data, prev: nil, next: nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *DoublyLinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	dll := DoublyLinkedList{}
	dll.append(1)
	dll.append(2)
	dll.append(3)

	fmt.Println("Doubly Linked List:")
	dll.display()
}
