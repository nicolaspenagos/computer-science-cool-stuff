package main

import "fmt"

type BST struct {
	Root *Node
}

type Node struct {
	Data  int64
	Left  *Node
	Right *Node
}

func (bst *BST) Insert(data int64) {
	bst.Root = insertNode(bst.Root, data)
}

func (bst *BST) Search(data int64) bool {
	return searchNode(bst.Root, data)
}

func (bst *BST) PrintInOrderTraversal() {
	inOrder(bst.Root)
	fmt.Println("")
}

func searchNode(node *Node, data int64) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	if data < node.Data {
		return searchNode(node.Left, data)
	}

	return searchNode(node.Right, data)
}

func insertNode(node *Node, data int64) *Node {
	if node == nil {
		return &Node{Data: data}
	}

	if node.Data > data {
		node.Left = insertNode(node.Left, data)
	} else {
		node.Right = insertNode(node.Right, data)
	}

	return node
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Print(node.Data, " ")
		inOrder(node.Right)
	}
}

func main() {
	bst := &BST{}

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)

	bst.PrintInOrderTraversal()

	fmt.Println(bst.Search(4))
	fmt.Println(bst.Search(99))
}
