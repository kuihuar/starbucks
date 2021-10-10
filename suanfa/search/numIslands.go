package search

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	n := len(grid)
	m := len(grid[0])
	fmt.Println("行=%d", n)
	fmt.Println("列=%d", m)
	if n == 0 {
		return 0
	}
	var dfsMarking func([][]byte, int, int)
	dfsMarking = func(grid [][]byte, i int, j int) {
		if i < 0 || j < 0 || i >=n || j >= m || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		dfsMarking(grid,i, j+1)
		dfsMarking(grid,i, j-1)
		dfsMarking(grid,i-1, j)
		dfsMarking(grid,i+1, j)
	}


	for i:=0; i<n;i++{
		for j:=0;j<m;j++{
			fmt.Printf("i=%d\ti=%d\t, grid[i][j]=%c\n",i,j,grid[i][j])
			if grid[i][j] == '1'{
				count++
				dfsMarking(grid,i,j)
			}
		}
	}
	fmt.Println(grid)
	return count
}

func TestnumIslands()  {
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
