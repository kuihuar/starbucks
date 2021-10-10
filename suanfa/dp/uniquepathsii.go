package dp

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i:=0; i< len(obstacleGrid);i++{
		for j:=0; j<len(obstacleGrid[0]);j++{
				if obstacleGrid[i][j] == 1 {// 障碍物直接标 0
					obstacleGrid[i][j] = 0
					// 首个位置不为障碍物时需要置位1（仅为了逻辑上的一致性，其实与实际不符；实际上应该是紧邻的两个位置为1）
					// 因为要从这里开始循环 （if j > 0 {obstacleGrid[i][j] += obstacleGrid[i][j-1]}这条语句
				}else if i==0 && j==0{
					obstacleGrid[i][j] = 1
				}else{
					obstacleGrid[i][j] = 0
					// 其他情况均为 左、上之和
					//好多的坑啊，这两个IF语句不能换位置
					if j > 0 {//先累加左边的
						obstacleGrid[i][j] += obstacleGrid[i][j-1]
					}
					//先累加上边的，每行第一个元素
					if i > 0 {
						obstacleGrid[i][j] += obstacleGrid[i-1][j]
					}
				}
		}
	}
	fmt.Println(obstacleGrid)
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m:= len(obstacleGrid)
	n:= len(obstacleGrid[0])
	dp := make([]int,n+1)
	dp[1] = 1
	for i:=0; i<m;i++{
		for j:=1;j<=n;j++{
			if obstacleGrid[i][j-1] == 1 {
				dp[j] = 0
			}else{
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[n]
}
func TestUniquePathsii()  {
	param := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	res := uniquePathsWithObstacles1(param)
	fmt.Println(res)
}