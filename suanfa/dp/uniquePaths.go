package dp

import "fmt"
//从start 走到（i,j）的不同路径数
func uniquePaths(m int, n int) int {
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	for i:=0; i< m;i++{
		dp[i][0] = 1
	}
	for i:=0; i<n;i++{
		dp[0][i] = 1
	}
	fmt.Println("%+v",dp)
	for i:=1; i<m;i++{
		for j:=1; j<n;j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			fmt.Println("%+v",dp)
		}
	}
	return dp[m-1][n-1]
}

//从（i,j）走到end的不同路径数
func uniquePaths1(m int, n int) int {
	dp := make([][]int,m)
	for i:=0; i<m ;i++{
		dp[i] = make([]int,n)
	}
	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if j==m-1 || j== n-1 {
				dp[i][j] = 1
			}else{
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
func TestUniquePaths()  {
	res := uniquePaths(3,7)
	fmt.Println(res)
}