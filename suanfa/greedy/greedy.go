package greedy

import "sort"

//贪心算法: 当下做最优判断，不能回退
//回溯： 能够回退
//动态规划： 最优判断 + 回退
//贪心算法解决一些最优问题：最小生成树，哈夫曼
//
////贪心算法，最大选项能够整除后面的选，能用1个20块，就不用两个10块
////如果是 10，9，1， 需要拼 18的话，如果先选10，就不是最优了
///贪心算法，问题能够分解成子问题，子问题的最优解能够递推一到最终问题的最优解
////贪心心算法和动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。
////动态规划则会保存以前的运算结果，并根据以前的结果对当前做出先择。有回退功能。


func findContentChildren(g []int, s []int) int {
	var ans int
	sort.Ints(g)
	sort.Ints(s)
	n,m := len(g),len(s)

	for i,j:=0,0; i< n && j<m; i++{
		for j<m &&  g[i] > s[j] {
			j++
		}
		if j<m {
			ans++
			j++
		}
	}
	return ans
}