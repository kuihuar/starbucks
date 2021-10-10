package datastructures

import "fmt"

type bstNode struct{
	value int
	left *bstNode
	right *bstNode
}

type bst struct {
	root *bstNode
}

func (b *bst)reset()  {
	b.root = nil
}
func (b *bst)insert(value int) {
	b.insertRecursive(b.root, value)
}
func (b *bst)insertRecursive(node *bstNode, value int) *bstNode {
	if b.root == nil {
		b.root = &bstNode{value: value}
		return b.root
	}
	if node == nil {
		return &bstNode{value: value}
	}
	if value <= node.value {
		node.left = b.insertRecursive(node.left, value)
	}
	if value > node.value {
		node.right = b.insertRecursive(node.right, value)
	}
	return node
}
func (b *bst) find(value int) error {
	node := b.findRecurive(b.root, value)
	if node == nil {
		return fmt.Errorf("not found %d", value)
	}
	return nil
}
func (b *bst)findRecurive(node *bstNode, value int) *bstNode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return b.root
	}
	if value < node.value {
		return b.findRecurive(node.left, value)
	}
	return b.findRecurive(node.right, value)
}
func (b *bst)inorder()  {
	b.inorderRecursive(b.root)
}
func (b *bst)inorderRecursive(node *bstNode)  {
	if node != nil {
		b.inorderRecursive(node.left)
		fmt.Println(node.value)
		b.inorderRecursive(node.right)
	}
}