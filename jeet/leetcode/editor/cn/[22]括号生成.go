//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 有效括号组合需满足：左括号必须以正确的顺序闭合。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 
//
// 示例 2： 
//
// 
//输入：n = 1
//输出：["()"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 8 
// 
// Related Topics 字符串 动态规划 回溯 👍 2007 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	res := generate(0,0,n, "")
	return res
}
func generate(left, right int, n int, parenthes string) []string{
	var res []string
	if left ==n && right == n{
		res = append(res,parenthes)
		return
	}
	if left < n{
		generate(left+1, right, n, parenthes+"(")
	}
	if right <left {
		generate(left, right+1, n, parenthes+")")
	}
}
//leetcode submit region end(Prohibit modification and deletion)
