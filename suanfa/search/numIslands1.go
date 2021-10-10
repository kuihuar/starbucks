package search

import "fmt"

func numIslands1(grid [][]byte) int {
	count := 0

	n := len(grid)
	m := len(grid[0])
	if n == 0 {
		return 0
	}
	for i:=0; i<n;i++{
		for j:=0;j<m;j++{
			if grid[i][j] == '1'{
				dfsMarking1(grid,j,j, n,m)
				count++
			}
		}
	}
	return count
}
func dfsMarking1(grid [][]byte, i, j ,n, m int)  {
	if i < 0 || j < 0 || i >=n || j >= m || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfsMarking1(grid, i, j+1, n, m)
	dfsMarking1(grid, i, j-1, n, m)
	dfsMarking1(grid, i-1, j, n, m)
	dfsMarking1(grid, i+1, j, n, m)
}

func TestnumIslands1()  {
	//grid := [][]byte{
	//	{'1','1','0','0','0'},
	//	{'1','1','0','0','0'},
	//	{'0','0','1','0','0'},
	//	{'0','0','0','1','1'},
	//}
	//grid := [][]byte{
	//	{'1','1','1','1','0'},
	//	{'1','1','0','1','0'},
	//	{'1','1','0','0','0'},
	//	{'0','0','0','0','0'},
	//}
	grid := [][]byte{
		{'1','0','1','1','0','1','1'},
	}
	res := numIslands(grid)
	fmt.Println(res)
}
