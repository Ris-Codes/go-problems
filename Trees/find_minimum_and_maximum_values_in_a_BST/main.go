package main

import "fmt"

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("The minimum value is: ", FindMin(root))
	fmt.Println("The maximum value is: ", FindMax(root))
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) Insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

func FindMin(node *Node) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

func FindMax(node *Node) int {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current.Value
}
