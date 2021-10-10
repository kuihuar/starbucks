package dp

func longestCommonSubsequence(text1 string, text2 string) int {

	m:= len(text1)
	n:= len(text2)

	dp := make([][]int, m+1)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n+1)
	}
	for i:=1; i< m+1;i++{
		for j:=1;j<n+1;j++{
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			}else{
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] =dp[i-1][j]
				}else{
					dp[i][j] =dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}