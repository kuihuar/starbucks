package devideconquer

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	ans := make([][]int,0)
	temp := []int{}
	dfs(&ans, nums,temp, 0)
	return ans
}
//ans 结果，index，list当前先或不选,不能通过，下边的可以通过
func dfs(ans *[][]int, nums []int, list []int, index int)  {
	if  index == len(nums) {
		fmt.Println(list)
		*ans = append(*ans, list)
		return
	}
	list= append(list,nums[index])
	dfs(ans, nums, list, index+1)
	//list随着递归一直在变动，不是本层的本地变量
	//这里改变会影响上面几层的函数，所以这里要去掉
	list = (list)[:len(list)-1]
	dfs(ans, nums, list, index+1)
}
func TestSubsets()  {
	param:= []int{9,0,3,5,7}
	//param:= []int{1,2,3}
	res := subsets2(param)
	fmt.Println(res)
	//fmt.Println(len(res))
}


func subsets1(nums []int) [][]int {
	set := []int{}
	var ans [][]int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, set...)) //什么意思 ,关键啊 why
			return
		}
		set = append(set, nums[cur]) //这里的顺序是啥意思
		dfs(cur+1)
		set = set[:len(set)-1]
		dfs(cur+1)
	}
	dfs(0)
	//fmt.Println(ans)
	return ans
}

//迭代
//往之前已经产生过的子集合里面加就可以。直到最后一个数加进去
func subsets2(nums []int) [][]int {
	result := [][]int{}
	result = append(result,[]int{})
	for _, num := range nums {
		for _,resultItem := range result {
			resultItemCopy := make([]int,len(resultItem))
			copy(resultItemCopy,resultItem)// 晕死，这里如果不copy的话，总是不成功
			result = append(result, append(resultItemCopy,num))
		}
	}
	return result
}