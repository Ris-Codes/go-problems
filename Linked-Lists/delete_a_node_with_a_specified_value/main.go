package main

import "fmt"

func main() {
	list := &LinkedList{}
	list.append(10)
	list.append(20)
	list.append(30)
	list.display()

	list.delete(20)
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

func (list *LinkedList) delete(value int){
	if list.head == nil {
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil && current.next.data != value{
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}