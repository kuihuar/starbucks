package bst

type TreeNode struct {
     Val   int
     Left  *TreeNode
     Right *TreeNode
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root,q,p)
}
func dfs(node, p, q *TreeNode) *TreeNode {
	if node.Val > p.Val && node.Val > q.Val {
		return dfs(node.Right, p,q)
	}else if node.Val < p.Val && node.Val < q.Val {
		return dfs(node.Left,p, q)
	}else{
		return node
	}
}

