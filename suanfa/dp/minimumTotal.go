package dp

import "fmt"

//DP
func minimumTotal(triangle [][]int) int {
	for i:=len(triangle)-2;i>=0;i--{
		for j:=0; j<len(triangle[i]);j++{
			//下边或者右边
			if triangle[i+1][j] > triangle[i+1][j+1]{
				triangle[i][j] += triangle[i+1][j+1]
			}else{
				triangle[i][j] += triangle[i+1][j]
			}
		}
	}
	return triangle[0][0]
}
//新开一维数组
func minimumTotalDp(triangle [][]int) int {
	res := triangle[len(triangle)-1]//最后一行
	for i:=len(triangle)-2;i>=0;i--{
		for j:=0; j<len(triangle[i]);j++{
			//下边或者右边
			if triangle[i+1][j] > triangle[i+1][j+1]{
				res[j] += triangle[i+1][j+1]
			}else{
				res[j] += triangle[i+1][j]
			}
		}
	}
	return res[0]
}

//recursion
func minimumTotal1(triangle [][]int) int {
	var helper func(int, int) int
	row := len(triangle)
	helper = func(level, c int) int {
		if level == row -1{
			return triangle[level][c]
		}
		left := helper(level+1, c)
		right := helper(level+1, c+1)
		if left > right {
			return right + triangle[level][c]
		}else{
			return left + triangle[level][c]
		}
	}
	return helper(0,0)
}

func minimumTotal2(triangle [][]int) int {
	var helper func(int, int) int
	mem := make([][]int,len(triangle))
	for i:=0;i<len(triangle);i++{
		mem[i] = make([]int,len(triangle[0]))
	}
	row := len(triangle)
	helper = func(level, c int) int {

		if mem[level][c] != 0 {
			return mem[level][c]
		}

		if level == row -1{
			mem[level][c] = triangle[level][c]
			return triangle[level][c]
		}
		left := helper(level+1, c)
		right := helper(level+1, c+1)
		if left > right {
			triangle[level][c] =  right + triangle[level][c]
		}else{
			triangle[level][c]=  left + triangle[level][c]
		}
		return triangle[level][c]
	}
	return helper(0,0)
}
func TestminimumTotal ()  {
	p := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	res := minimumTotal1(p)
	fmt.Println(res)

}
