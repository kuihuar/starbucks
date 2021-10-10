package recircution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
//递归
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <=lower || root.Val>=upper {
		return false
	}
	//调用左子树的时候，把上界upper改为root.val,因为左子树里所有节点的值都小于根节点的值
	//调用右子树的时候，把下界lower改为root.val
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
