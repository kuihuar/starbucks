package stack_style

import "fmt"

//1. 暴力

//func largestRectangleArea(heights []int) int {
//	var maxArea int
//	var minHeight int
//	if len(heights)== 1 {
//		return heights[0]
//	}
//	for i:=0; i< len(heights); i++ {
//		minHeight = heights[i]
//		for j:=i; j< len(heights);j++{
//			if heights[j] < minHeight {
//				minHeight = heights[j]
//			}
//			area := (j-i+1) * minHeight
//			if area > maxArea {
//				maxArea = area
//			}
//			fmt.Printf("i=%d\tj=%d\tmaxarea=%d\n",i,j,maxArea)
//		}
//	}
//	return maxArea
//}
// 暴力2
// 固定中间柱形的高度，确定左边界和右边界
// area = height[i] = (right-left)
//func largestRectangleArea(heights []int) int {
//	var maxArea int = 0
//	for i:=0; i< len(heights); i++ {
//		minHeight := heights[i]
//		for j := i; j< len(heights);j++ {
//			currHeight := heights[j]
//			if currHeight < minHeight {
//				minHeight = currHeight
//			}
//			area := minHeight * (j-i+1)
//			//fmt.Printf("area=%d := %d * (%d-%d+1)\t",area,minHeight,j,i)
//			if area > maxArea {
//				maxArea = area
//			}
//			//fmt.Printf("maxArea=%d\n",maxArea)
//		}
//	}
//	return maxArea
//}
// 栈
// 如果加入的元素比栈顶元素大，则无法确定右边界
// 如果要加入的元素比栈顶元素小，则可以确定右边界 rightBound，即是当前元素的大小
func largestRectangleArea(heights []int) int {
	var maxArea int

	stack := []int{-1} // -1 是下标
	var rightBound int
	for i:=0; i< len(heights); i++ {
		rightBound = i
		for stack[len(stack) -1] != -1 && heights[stack[len(stack) -1]] >= heights[i] { //发现栈顶元素大于要入栈的
			peek := stack[len(stack) -1] //这两句是POP
			stack = stack[:len(stack)-1]
			area :=  heights[peek] * (rightBound -  stack[len(stack) -1] - 1)
			if area > maxArea {
				maxArea = area
			}
			fmt.Printf("%d*(%d-(%d)-1)\t",heights[peek],rightBound,stack[len(stack) -1])
			fmt.Printf("area=%d\tmaxArea=%d\n",area,maxArea)
		}
		stack = append(stack,i)
	}
	for stack[len(stack) -1] != -1{
		peek := stack[len(stack) -1]
		stack = stack[:len(stack)-1]
		area :=  heights[peek] * (rightBound -  stack[len(stack) -1])
		if area > maxArea {
			maxArea = area
		}
		fmt.Printf("%d*(%d-(%d)-1)\t",heights[peek],rightBound,stack[len(stack) -1])
		fmt.Printf("area=%d\tmaxArea=%d\n",area,maxArea)
	}
	return maxArea
}
func TestLRA() int {
	//heights := []int{6,7,5,2,4,5,9,3}
	heights := []int{2,3}
	return largestRectangleArea(heights)
}

