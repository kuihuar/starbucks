package sudo
//
//import "fmt"
//
//func solveSudoku(board [][]byte)  {
//	if board == nil || len(board)==0 {
//		return
//	}
//	dfs(board)
//}
//func dfs(board [][]byte) bool{
//	// 遍历整个棋盘
//	for i:=0; i<len(board);i++{
//		for j:=0;j<len(board[0]);j++{
//			if board[i][j] == '.' { //棋盘为空位
//				for c :='1'; c<='9';c++{ // 尝试填 1-9
//					if isvalid(board, i,j,c) { //整个棋盘是否合法
//						board[i][j] = byte(c)
//						if dfs(board) { //整个棋盘已经解决
//							return true
//						}
//						//回溯
//						board[i][j] = '.'
//					}
//				}
//				//如果什么都不行就返回false
//				return false
//			}
//		}
//	}
//	//整个棋盘搞完
//	return true
//}
//func isvalid(board [][]byte, row, col int, c int32) bool {
//	for i:=0; i<9;i++{
//		if board[i][col] != '.' && board[i][col] == byte(c) {
//			return false
//		}
//		if board[row][i] != '.' && board[row][i] == byte(c) {
//			return  false
//		}
//		//根据 boxIndex 转换行和列
//		boxXindex :=  (row / 3) * 3 + i / 3
//		boxYindex :=  (col / 3) * 3 + i % 3
//		if board[boxXindex][boxYindex] != '.' && board[boxXindex][boxYindex] == byte(c) {
//			return false
//		}
//	}
//	return true
//}
//func TestSolveSuduku()  {
//	param := [][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},
//		{'.','9','8','.','.','.','.','6','.'}, {'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},
//		{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},
//		{'.','.','.','.','8','.','.','7','9'}}
//	solveSudoku(param)
//	//fmt.Println(param)
//	for i:=0; i<9;i++{
//		for j:=0; j<9;j++{
//			fmt.Printf("%s  ",string(param[i][j]))
//		}
//		fmt.Println()
//	}
//}