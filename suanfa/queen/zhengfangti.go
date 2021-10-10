package queen

import (
	"fmt"
)

func TestCicle()  {
	n:=4
	matrix := make([][]string,n)

	for i:=0;i<n;i++{
		matrix[i]= make([]string,n+1)
	}

	left,top :=0,0
	right:=len(matrix[0])-1
	button :=len(matrix)-1

	for{
		for i:=left; i<=right;i++{
			fmt.Printf("%d-%d\t",top,i)
		}
		top++
		if top > button {
			break
		}
		for i:=top; i<=button;i++{
			fmt.Printf("%d-%d\t",i,right)
		}
		right--
		if left > right{
			//break
		}
		for i:= right;i>=left;i--{
			fmt.Printf("%d-%d\t",button,i)
		}
		button--
		if top > button{
			//break
		}
		for i:= button;i>=top;i--{
			fmt.Printf("%d-%d\t",i,left)
		}
		left++
		if left > right {
			//break
		}
		//return
		fmt.Println("aaa")
	}
}
//class Solution {
//	public List<Integer> spiralOrder(int[][] matrix) {
//
//		List<Integer> arr = new ArrayList<>();
//		int left = 0, right = matrix[0].length-1;
//		int top = 0, down = matrix.length-1;
//
//		while (true) {
//			for (int i = left; i <= right; ++i) {
//				arr.add(matrix[top][i]);
//			}
//			top++;
//			if (top > down) break;
//			for (int i = top; i <= down; ++i) {
//				arr.add(matrix[i][right]);
//			}
//			right--;
//			if (left > right) break;
//			for (int i = right; i >= left; --i) {
//				arr.add(matrix[down][i]);
//			}
//			down--;
//			if (top > down) break;
//			for (int i = down; i >= top; --i) {
//				arr.add(matrix[i][left]);
//			}
//			left++;
//			if (left > right) break;
//
//			}
//			return arr;
//		}
//}
