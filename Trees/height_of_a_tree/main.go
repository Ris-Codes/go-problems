package main

import "fmt"

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	fmt.Println("Height of the Tree is: ", Height(root))
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func Height(node *Node) int {
	if node == nil {
		return 0
	}
	leftHeight := Height(node.Left)
	rightHeight := Height(node.Right)
	return 1 + max(leftHeight, rightHeight)
}
