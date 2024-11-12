package main

import "fmt"

func main() {
	root := &BSTNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Print("Pre Order Traversal: ")
	root.PreOrder()
	fmt.Println()

	fmt.Print("In Order Traversal: ")
	root.InOrder()
	fmt.Println()

	fmt.Print("Post Order Traversal: ")
	root.PostOrder()
	fmt.Println()

	fmt.Println("Search 7: ", root.Search(7))
	fmt.Println("Search 20: ", root.Search(20))
}

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

// Inserting
func (node *BSTNode) Insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BSTNode{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BSTNode{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

// Searching
func (node *BSTNode) Search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.Search(value)
	}
	if value > node.Value {
		return node.Right.Search(value)
	}
	return true
}

// Traversals in BST
// -------------1. PreOrder Traversal
func (node *BSTNode) PreOrder() {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	node.Left.PreOrder()
	node.Right.PreOrder()
}

// ---------------2. InOrder Traversal
func (node *BSTNode) InOrder() {
	if node == nil {
		return
	}
	node.Left.InOrder()
	fmt.Printf("%d ", node.Value)
	node.Right.InOrder()
}

// ---------------3. PostOrder Traversal
func (node *BSTNode) PostOrder() {
	if node == nil {
		return
	}
	node.Left.PostOrder()
	node.Right.PostOrder()
	fmt.Printf("%d ", node.Value)
}