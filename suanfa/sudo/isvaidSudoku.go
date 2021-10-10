package sudo

import (
	"fmt"
)

//行中无重复
//列中无重复
//3*3 box无重复
// 判断某行某列的某个数字第二次出现，将board[i][j]转成数字0-8
func isValidSudoku(board [][]byte) bool {
	var rows, columns,boxes [9][9]int
	for i := 0; i < 9;i++{
		for j := 0; j < 9;j++{
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'
			rows[i][num]++
			columns[j][num]++
			boxIndex := (i / 3) * 3 + j / 3
			boxes[boxIndex][num]++
			if rows[i][num] > 1 || columns[j][num] > 1 || boxes[boxIndex][num] > 1 {
				return false
			}
		}
	}
	return true
}

func Testsudo (){
	board := [][]byte{{'8','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}
	// {'8','3','.','.','7','.','.','.','.'},
	// {'6','.','.','1','9','5','.','.','.'},
	// {'.','9','8','.','.','.','.','6','.'},
	// {'8','.','.','.','6','.','.','.','3'},
	// {'4','.','.','8','.','3','.','.','1'},
	// {'7','.','.','.','2','.','.','.','6'},
	// {'.','6','.','.','.','.','2','8','.'},
	// {'.','.','.','4','1','9','.','.','5'},
	// {'.','.','.','.','8','.','.','7','9'}}
	res := isValidSudoku(board)
	fmt.Println(res)
}