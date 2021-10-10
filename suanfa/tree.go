package suanfa

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Constructor() TreeNode {
	return TreeNode{}
}
func NewTreeNode() *TreeNode {
	return &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
}
//root left right
func (root *TreeNode)PreOrder()  {
	fmt.Println(root.Val)
	root.Left.PreOrder()
	root.Right.PreOrder()
}
//left-root-right
func (root *TreeNode)InOrder(){
	root.Left.InOrder()
	fmt.Println(root.Val)
	root.Left.InOrder()
}
//left-right-root
func (root *TreeNode)PostOrder(){
	root.Left.PostOrder()
	root.Right.PostOrder()
	fmt.Println(root.Val)
}
//递归
func PreorderTraversal(root *TreeNode)(vals []int){
	var preorder func(node *TreeNode)
	//新定义一个方法去递归，类似于helper函数
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals,node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	return
}
//迭代(模拟的递归函数的栈)
func PreorderTraversal_1(root *TreeNode)(vals []int){
	stack := []*TreeNode{}
	node := root
	for node != nil && len(stack) > 0 {
		for node != nil {
			vals = append(vals,node.Val)
			stack = append(stack,node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}