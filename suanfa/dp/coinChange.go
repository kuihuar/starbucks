package dp

import (
	"fmt"
	"math"
)

//每层表示用掉一个硬币
//暴力，每次可以取1，取2，也可以取5，和就是要sum
//把所有不同的路径求出来
// 人肉递归，状态树，
//从11开始，减去1个1，求布值10有多少不同的走法，或者是先减了1个2，就变成9这个子问题，或者是减去1个5, 变成凑成6的子问题
// 叶子节点是到0表示刚好凑对了，得到一组解，把这个结果往上传，最后得到最优的（硬币数最少的情况）,把最少的情况比较出来。
// 叶子节点是负的表示凑不到，非法的，停掉即可。不需要再走了
//树的深度表示用的硬币的个数，深度层交为1（10，9，6）表示只用了一个硬币，
//村的边表示用的硬币的大小
//问题变成这棵树上找到结点为0的结点，且层次是最低的（深度要最小）
//广度优先遍历，直到第一次碰到数值为0的节点，当前的层就是我们要的答案

//递归，指数
//BSF


//这个题目看leetcode,好多种解法
//DP f(n) = min{f(n-k), for k in [1,2,5]} + 1
func coinChange(coins []int, amount int) int {
	max := amount +1 //最大金额
	dp := make([]int,max)//数组长度最大为max,也就是每次兑换1块，即需要长度为 Max
	for i:=0;i<max;i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0 //0级台阶，最少走0步
	var min func(int, int) int
	min = func(i int, i2 int) int {
		if i < i2 {
			return i
		}else{
			return i2
		}
	}
	for i:=1;i<=amount;i++ { //之前的步数
		for j:=0; j<len(coins);j++{
			if coins[j] <= i { //要凑的面值数i
				// dp[i]
				// coins[j], 当前用的
				// i-coins[j] 剩下的面值
				// dp[i-coins[j]] +1 剩下的面值 + 1
				dp[i] = min(dp[i], dp[i-coins[j]] +1 )
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func TestCoinChange()  {
	res := coinChange([]int{1,2,5}, 11)
	fmt.Println(res)
}