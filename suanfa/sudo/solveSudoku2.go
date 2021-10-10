package sudo

import "fmt"

func solveSudoku1(board [][]byte)  {
	if board == nil || len(board)==0 {
		return
	}
	dfs1(board,0)
}
func dfs1(board [][]byte, level int) bool{
	if level == 81 {
		return true
	}
	i,j := level / 9, level % 9
	if board[i][j] != '.' { //当前位置已是数字，往下一层
		return dfs1(board, level+1)
	}
	flag := [10]bool{}
	validate(board,i,j,&flag)
	for k:=1; k<=9;k++{
		if flag[k] {
			board[i][j] = byte('0' +k)
			//fmt.Printf("%s",string(board[i][j]))
			if dfs1(board, level+1) {
				return true
			}
		}
	}
	board[i][j] = '.' //回溯
	return false
}
//存储1-9是否有重复
func validate(board [][]byte, row, col int, flag *[10]bool) {
	for k:=0;k<len(flag);k++{
		flag[k] = true
	}
	for i:=0; i<9;i++{
		if board[row][i] != '.'  {
			flag[board[row][i]-'0'] = false
		}
		if board[i][col] != '.' {
			flag[board[i][col]-'0'] = false
		}
		//根据 boxIndex 转换行和列
		boxXindex :=  (row / 3) * 3 + i / 3
		boxYindex :=  (col / 3) * 3 + i % 3
		if board[boxXindex][boxYindex] != '.' {
			flag[board[boxXindex][boxYindex] -'0'] = false
		}
	}
}
func TestSolveSuduku1()  {
	param := [][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'}, {'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}}
	solveSudoku1(param)
	//fmt.Println(param)
	for i:=0; i<9;i++{
		for j:=0; j<9;j++{
			fmt.Printf("%s  ",string(param[i][j]))
		}
		fmt.Println()
	}
}