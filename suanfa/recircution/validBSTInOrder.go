package recircution

import "math"


// 中序带一个参数，先序带两个参数

func isValidBST1(root *TreeNode) bool {
	min := math.MinInt64
	return inorder(root, &min)
}
func inorder(node *TreeNode, last *int) bool{

	if node== nil {
		return true
	}
	l := inorder(node.Left, last)
	if *last >= node.Val{
		return false
	}
	*last = node.Val
	r := inorder(node.Right, last)
	return l && r
}


//先序
func preorder(node *TreeNode, prev, next *TreeNode) bool {
	if node == nil {
		return true
	}
	if prev != nil && prev.Val>= node.Val {
		return false
	}
	if next != nil && next.Val <= node.Val {
		return false
	}
	return preorder(node.Left, prev, node) && preorder(node.Right, node, next)
}
func TisValidBST1()  {
	root := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	isValidBST1(root)
}
//递归

