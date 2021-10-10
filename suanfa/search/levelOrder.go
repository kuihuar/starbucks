package search

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 1.bfs
// 2.
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return  res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		child := []int{}
		tempQueue := []*TreeNode{}
		for i:=0; i< len(queue);i++{
			node := queue[i]
			if node.Left!= nil {
				tempQueue =append(tempQueue,node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
			child = append(child, node.Val)
		}
		res = append(res, child)
		queue = tempQueue
	}
	return res
}
