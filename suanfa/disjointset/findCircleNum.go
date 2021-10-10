package disjointset

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := newUnionFind(n)
	//fmt.Println("uf=",uf)
	for i:=0; i<n-1; i++{
		for j:=i+1;j<n;j++{
			if isConnected[i][j] == 1{
				uf.union(i,j)
			}
		}
	}
	//fmt.Println("uf=",uf)
	//temp := make(map[int]bool)
	//for i:=0; i<n;i++{
	//	temp[ uf.fineParent(i)] = true
	//}
	//return len(temp)
	return uf.Count()
}

func TestFindCircleNum()  {
	param := [][]int{{1,1,0},{1,1,0},{0,0,1}}
	res := findCircleNum(param)
	fmt.Println(res)
}