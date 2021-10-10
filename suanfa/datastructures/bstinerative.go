package datastructures

import "fmt"

func (b *bst)insert1(value int)  {
	b.insertIterative(b.root,value)
}
func (b *bst)insertIterative(node *bstNode,value int)  {
	if b.root == nil {
		b.root = &bstNode{value: value}
	}
	if node == nil {
		return
	}
	var terminalNode *bstNode
	for node != nil {
		terminalNode = node
		if value <= node.value{
			node = node.left
		}else{
			node = node.right
		}
	}
	if value <= terminalNode.value {
		terminalNode.left = &bstNode{value: value}
	}else{
		terminalNode.right = &bstNode{value: value}
	}
	return
}
func (b *bst) find1(value int) error  {
	node := b.findIterator(b.root, value)
	if node == nil {
		return fmt.Errorf("not found value %d",value)
	}
	return nil
}
func (b *bst) findIterator(node *bstNode,value int) *bstNode {
	if node == nil {
		return nil
	}
	for node != nil {
		if value < node.value{
			node = node.left
		}else if value > node.value {
			node = node.right
		}else{
			return node
		}
	}
	return nil
}