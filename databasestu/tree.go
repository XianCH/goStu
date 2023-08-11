package main

import (
	"fmt"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(key int) {
	if bt.Root == nil {
		bt.Root = &Node{Key: key}
	} else {
		bt.insertRecursively(bt.Root, key)
	}
}

func (bt *BinaryTree) insertRecursively(node *Node, key int) {
	if key < node.Key {
		if node.Left == nil {
			node.Left = &Node{Key: key}
		} else {
			bt.insertRecursively(node.Left, key)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Key: key}
		} else {
			bt.insertRecursively(node.Right, key)
		}
	}
}

func (bt *BinaryTree) Search(key int) bool {
	return bt.searchRecursively(bt.Root, key)
}

func (bt *BinaryTree) searchRecursively(node *Node, key int) bool {
	if node == nil {
		return false
	}

	if key == node.Key {
		return true
	} else if key < node.Key {
		return bt.searchRecursively(node.Left, key)
	} else {
		return bt.searchRecursively(node.Right, key)
	}
}

func main() {
	binaryTree := BinaryTree{}
	keys := []int{5, 3, 8, 1, 4, 7, 9}

	for _, key := range keys {
		binaryTree.Insert(key)
	}

	fmt.Println("Binary Tree:")
	fmt.Println("In-order traversal:")
	binaryTree.InOrderTraversal(binaryTree.Root)
	fmt.Println()

	keyToSearch := 7
	fmt.Printf("Search for key %d: %v\n", keyToSearch, binaryTree.Search(keyToSearch))
}

func (bt *BinaryTree) InOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	bt.InOrderTraversal(node.Left)
	fmt.Printf("%d ", node.Key)
	bt.InOrderTraversal(node.Right)
}
