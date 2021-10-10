package queen

import "fmt"

var solutions [][]string
func solveNQueens(n int) [][]string {
	//solutions = [][]string{}
	queens := make([]int,n)
	for i:=0;i<n;i++{
		queens[i] = -1
	}
	columns := map[int]bool{}
	main := map[int]bool{}
	sub := map[int]bool{}
	dfs(queens, n, 0, columns, main, sub)
	return solutions
}

func dfs(queens []int, n,row int, columns, main, sub map[int]bool)  {
	if row == n {
		board := generateBoard(queens,n)
		solutions = append(solutions,board)
		return
	}
	// 1. 同一列 columns[i]
	// 2. 捺主对角线横坐标减纵坐标（看左上角是否有Q）
	// 3. 撇副对角线横坐标加纵坐值相同（看右上角的是否有Q）
	for col:=0; col< n;col++{
		if columns[col] || main[row-col] || sub[row+col] {
			continue
		}
		queens[row]=col//第row行放置Queue
		//fmt.Println("queens=",queens)
		columns[col],main[row-col],sub[row+col]=true,true,true
		dfs(queens,n, row+1, columns,main,sub)
		//queens[row] = -1
		delete(columns,col)
		delete(main,row-col)
		delete(sub,row+col)
	}
}
func generateBoard(queens []int,n int) []string {
	//fmt.Println("queens=",queens)
	board:= []string{}
	for i:=0;i<n;i++{
		row := make([]byte,n)
		for j:=0;j<n;j++{
			row[j] = '.'
		}
		row[queens[i]]= 'Q'
		board = append(board,string(row))
	}
	return board
}
func TestNqueens()  {
	n:= 4
	res := solveNQueens(n)
	for i:=0;i<len(res);i++{
		for j:=0;j<n;j++{
			fmt.Printf("%s\n",res[i][j])
		}
		fmt.Println()
	}
	fmt.Println(res)
}