package main

import "fmt"

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	fmt.Println("Number of leaf nodes in the binary tree: ", CountLeafNodes(root))
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func CountLeafNodes(node *Node) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	return CountLeafNodes(node.Left) + CountLeafNodes(node.Right)
}
