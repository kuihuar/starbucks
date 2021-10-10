package dfs

import "fmt"

//朋友圈数量


func findCircleNum(isConnected [][]int) int {
	visited := make([]bool,len(isConnected))
	res := 0
	for i:=0; i<len(isConnected);i++ {
		fmt.Printf("i=%d\n",i)
		if !visited[i] {
			dfs(isConnected, visited, i)
			res++
		}
	}
	return  res
}
func dfs(isConnected [][]int, visited []bool, i int)  {
	for j:=0;j< len(isConnected);j++{
		fmt.Printf("isConnected[%d][%d]\n",i,j)
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(isConnected, visited,j)
		}
	}
}

func TestFindCircleNum()  {
	param := [][]int{{1,1,0},{1,1,0},{0,0,1}}
	ret := findCircleNum(param)
	fmt.Println(ret)
}