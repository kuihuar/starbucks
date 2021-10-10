//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x⁴
//
// Related Topics 递归 数学 👍 721 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if x == 1 || n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / helper(x,-n)
	}else{
		return helper(x,n)
	}
}
func helper(x float64, n int) float64  {
	if n == 1{
		return x
	}
	half := myPow(x, n / 2)
	if n % 2 == 0 {
		return half * half
	}else {
		return half * half * x
	}
}
//leetcode submit region end(Prohibit modification and deletion)
