package tree

import "fmt"

//bfs
func jump1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	queue := []int{0}
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for i :=0; i<size;i++{
			jumpIndex := queue[0]//拿一个值
			queue = queue[1:]
			if jumpIndex == len(nums) -1 {
				return level
			}
			for j:=1;j<= nums[jumpIndex];j++{//往前跳
				queue= append(queue, jumpIndex +j)
			}
		}
		level++
	}
	return 0
}

//贪心算法，
// 贪心策略
func jump2(nums []int) int {

	steps := 0
	start := 0
	end:= 0
	//end 一直走到nums， 或超过 nums
	for end < len(nums) - 1 {
		//每次从start 到 end 区间，能走到最远的位置
		maxPos := 0
		for i:= start; i<=end;i++{
			// nums[i] 可以跳多远， i 表示当前的位置
			if i + nums[i] > maxPos{
				maxPos = i + nums[i]
			}
		}
		start = end +1
		end = maxPos
		steps++
	}
	return steps
}
func jump3(nums []int) int {

	steps := 0
	maxPos := 0
	end:= 0
	//end 一直走到nums， 或超过 nums
	for i:=0; i < len(nums) -1;i++ {
		if i + nums[i] > maxPos {
			maxPos= i + nums[i]
		}
		if i == end {
			steps++
			end = maxPos
		}
	}
	return steps
}

func jump4(nums []int) int{
	steps:=0 //总共需要跳的步数
	end:=0 //当前跳的最位置
	maxPos :=0
	for i:=0; i< len(nums)-1;i++{
		if i + nums[i] > maxPos{
			maxPos = i+ nums[i]
		}
		if i==end {
			steps++
			end = maxPos
		}
	}
	return steps
}
func TestBFS(){
	res := jump3([]int{2,3,1,1,4})
	fmt.Println(res)
}