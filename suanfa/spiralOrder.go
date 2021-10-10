package suanfa

//螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	left, top := 0, 0
	right, bottom := len(matrix[0]) -1, len(matrix) -1
	res := []int{}
	for {
		for i:=left;i<=right;i++{
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom{
			break
		}
		for i:=top;i<=bottom;i++{
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i:= right;i>=left;i--{
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		for i:= bottom; i>=top;i--{
			res = append(res, matrix[i][left])
		}
		left++
		if left > right{
			break
		}
	}
	return  res
}
