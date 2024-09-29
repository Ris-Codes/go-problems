package main

import "fmt"

func main() {
	list := &LinkedList{}
	list.append(10)
	list.append(20)
	list.append(30)
	list.display()

	list.insertAfter(20, 25)
	list.display()
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) append(data int) {
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
}

func (list *LinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList) insertAfter(x, data int) {
	current := list.head
	for current != nil {
		if current.data == x {
			newNode := &Node{data: data}
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}