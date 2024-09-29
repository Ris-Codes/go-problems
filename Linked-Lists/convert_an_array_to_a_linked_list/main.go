package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
    list := arrayToLinkedList(arr)
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

func arrayToLinkedList(arr []int) *LinkedList {
	list := &LinkedList{}
	for _, value := range arr {
		list.append(value)
	}
	return list
}
