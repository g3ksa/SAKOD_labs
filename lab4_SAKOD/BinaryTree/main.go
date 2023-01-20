package binaryTree

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value int
}

type BinaryTree struct {
	Root *BinaryNode
}

func BuildBalancedTree(data []int, min, max int) *BinaryTree {
	median := min + (max-min)/2
	return &BinaryTree{
		Root: &BinaryNode{
			Value: data[median],
			Left:  BuildNode(data, min, median),
			Right: BuildNode(data, median+1, max),
		},
	}
}

func BuildNode(data []int, min, max int) *BinaryNode {
	if min == max {
		return nil
	}

	median := min + (max-min)/2
	return &BinaryNode{
		Value: data[median],
		Left:  BuildNode(data, min, median),
		Right: BuildNode(data, median+1, max),
	}
}

func (t *BinaryTree) Insert(value int) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{Value: value, Left: nil, Right: nil}
	} else {
		t.Root.Insert(value)
	}
	return t
}

func (n *BinaryNode) Insert(value int) {
	if n == nil {
		return
	} else if value <= n.Value {
		if n.Left == nil {
			n.Left = &BinaryNode{Value: value, Left: nil, Right: nil}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryNode{Value: value, Left: nil, Right: nil}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (t *BinaryTree) PreOrder() {

}

func Print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Value)
	Print(w, node.Left, ns+2, 'L')
	Print(w, node.Right, ns+2, 'R')
}
