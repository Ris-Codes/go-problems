package main

import (
	"fmt"
	"math"
)

func main() {
	root := &Node{Value: 10}
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	min := math.MinInt
	max := math.MaxInt

	fmt.Println("Is this a valid BST? ", IsValidBST(root, min, max))
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

func IsValidBST(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Value <= min || node.Value >= max {
		return false
	}

	return IsValidBST(node.Left, min, node.Value) && IsValidBST(node.Right, node.Value, max)
}
