package main

import "fmt"

func main() {
	list := &DoublyLinkedList{}
    list.append(10)
    list.append(20)
    list.append(30)
    list.display()
}

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (list *DoublyLinkedList) append(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}

func (list *DoublyLinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}
