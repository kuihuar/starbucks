package str

import "fmt"
//还没有验证
func longestPalindrome(s string) string {
	n := len(s)

	res := ""
	dp := make([][]bool,n)
	for i:=0; i<n;i++{
		dp[i] = make([]bool,n)
	}
	//for i:=n; n >= 0; n-- {
	//	for j:=i;j<n;j++{
	//		if s[i] == s[j]  && (dp[i-1][j+1] || i-j < 2){
	//			dp[i][j] = true
	//		}
	//		if dp[i][j] && i-j+1 > len(res) {
	//			res = s[j:i+1]
	//		}
	//	}
	//}
	return res
}

func TestLongestPalindrome()  {
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}