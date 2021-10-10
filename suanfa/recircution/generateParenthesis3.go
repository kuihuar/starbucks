package recircution

import "fmt"

//func generateParenthesis(n int) []string {
//	generate(0, 0,n,"")
//	return nil
//}
//var in int
//func generate(left int, right int, total int, currentParenthes string)  {
//	//terminator
//	if left == total && right == total {
//		fmt.Println(currentParenthes)
//		return
//	}
//	//process current logic: 加左括号或加右括号
//	s1 := currentParenthes + "("
//	//drill down
//	if left < total {
//		generate(left +1, right,total, s1)
//	}
//	s2 := currentParenthes + ")"
//	if left > right{
//		generate(left, right+1,total, s2)
//	}
//	in++
//	fmt.Println(in)
//	//reverse states
//	//s1和s2都是局部变量，不用reverse去清除
//}



// i=n时相比n-1组括号增加的那一组括号的位置
// 当我们清楚所有i<n时括号的可能生成排列后，对i=n的情况，我们考虑整个括号排列中最左边的括号
// 它一定是一个左括号，那么它可以和它对应的右括号组成一组完整的括号"()",我们认为这一组是相比n-1增加进来的括号
// 那么剩下n-1组括号可能在哪呢？
// 剩下的括号要么在这一组新增的括号内部，要么 在这一组新增括号的外部（右侧）
// 既然知道了 i<n的情况，那我们就可以对所有情况进行遍历：
// "(" + [i=p时所有括号的排列组合] + ")" + [i=q时所有括号的排列组合]
// 其中p+q = n-1,且p,q均为非负整数
// 事实上，当上述p从0取到n-1,q从n-1取到0后，所有情况遍历完了



//  TODO DP方程，没有看懂
func generateParenthesis3(n int) []string {
	var res []string
	if n==0{
		return res
	}
	if n == 1{
		return []string{"()"}
	}
	dp := make([][]string,n+1)

	dp[0] =append(dp[0],"")
	dp[1] =append(dp[1],"()")
	//fmt.Println(dp)
	for i:=2; i<n+1;i++{
		for j:=0; j<i;j++{
			for p:=0;p<len(dp[j]);p++ {
				for q:=0; q<len(dp[i-j-1]);q++{
					fmt.Printf("i=%d,j=%d,p=%d,q=%d\t",i,j,p,q)
					fmt.Printf("dp前=%v\t",dp)
					fmt.Printf("dp[j][p]:dp[%d][%d]=%+v\t", j,p,dp[j][p])
					fmt.Printf("dp[i-j-1][q]:dp[%d-%d-1][%d]=%+v\t", i,j,q,dp[i-j-1][q])
					dp[i] =append(dp[i], "(" + dp[j][p] + ")" + dp[i-j-1][q])
					fmt.Printf("dp后=%v\n",dp)
					fmt.Println("++++++++++++++++++++++++")
				}
			}
		}
	}
	fmt.Printf("n=%d",n)
	return dp[n]
}


func TestGPDDP() []string {
	 res:= generateParenthesis3(3)
	 fmt.Println(res)
	 return res
}