package main

import "fmt"

func main() {
	root := &Node{Value: 10}
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Print("BST before Deletion: ")
	root.InOrder()
	fmt.Println()

	DeleteNode(root, 10)
	fmt.Print("BST after Deletion: ")
	root.InOrder()
	fmt.Println()
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

func (node *Node) InOrder() {
	if node == nil {
		return
	}
	node.Left.InOrder()
	fmt.Printf("%d ", node.Value)
	node.Right.InOrder()
}

func FindMin(node *Node) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

func DeleteNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.Value {
		node.Left = DeleteNode(node.Left, key)
	} else if key > node.Value {
		node.Right = DeleteNode(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		minValue := FindMin(node.Right)
		node.Value = minValue
		node.Right = DeleteNode(node.Right, minValue)
	}
	return node
}