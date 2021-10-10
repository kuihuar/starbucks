package tree

import (
	"fmt"
	"math"
)
var path []int
var minSteps int = math.MaxInt64
func jump(nums []int) int {
	var path []int
	dfs(nums, 0, path)
	if minSteps == math.MaxInt64 {
		return 0
	}
	return minSteps
}

func dfs(arr []int,jumpedIndex int, path []int)  {
	//fmt.Printf("jumpIndex=%d\t",jumpIndex)
	//fmt.Printf("minSteps=%d\t",minSteps)
	//fmt.Printf("arr=%v\t",arr)
	//fmt.Printf("arr=%v\n",path)
	if jumpedIndex == len(arr) -1 { //跳到最后一个位置
		if len(path)  < minSteps{
			minSteps = len(path)
		}
		return
	}

	for i:=1; i<= arr[jumpedIndex];i++ {
		if jumpedIndex + i >= len(arr) {
			continue
		}
		path = append(path,i)
		// jumpIndex+1 要跳到下一步的位置
		dfs(arr, jumpedIndex+i, path)
		path = path[:len(path)-1]
	}
}
func TestCanReach(){
	//res := jump([]int{2,3,1,1,4})
	res := jump([]int{5,6,4,4,6,9,4,4,7,4,4,8,2,6,8,1,5,9,6,5,2,7,9,7,9,6,9,4,1,6,8,8,4,4,2,0,3,8,5})
	fmt.Println(res)
}