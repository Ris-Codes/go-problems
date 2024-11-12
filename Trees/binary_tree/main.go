package main

import "fmt"

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	fmt.Print("Pre Order Traversal: ")
	PreOrder(root)
	fmt.Println()

	fmt.Print("In Order Traveresal: ")
	InOrder(root)
	fmt.Println()

	fmt.Print("Post Order Traversal: ")
	PostOrder(root)
	fmt.Println()
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Newnode
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// Traverssals through a Binary Tree
// -----------1. Pre Order Traversal (root -> left -> right)----------
func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// -----------2. In Order Traversal (left -> root -> right)----------
func InOrder(node *Node)  {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Printf("%d ", node.Value)
	InOrder(node.Right)
}

// -----------3. Post Order Traversal (left -> right -> root)----------
func PostOrder(node *Node)  {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Printf("%d ", node.Value)
}