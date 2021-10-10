package list

import "fmt"

//用双端队列实现
// 1. 暴力
//func maxSlidingWindow(nums []int, k int) []int {
//	var res []int
//	for i:=0; i<len(nums)-k+1;i++{
//		maxValue := max(nums[i:i+k])
//		res = append(res,maxValue)
//	}
//	return res
//}
//
//func max(intsLice []int) int {
//	var maxValue int = intsLice[0]
//	fmt.Println(intsLice)
//	for _,val := range intsLice{
//		if val > maxValue {
//			maxValue = val
//		}
//	}
//	return maxValue
//}
// 双向队列实现
func maxSlidingWindow(nums []int, k int) []int {
	fmt.Printf("%T,%v\n",nums, nums)
	queue := []int{}
	push := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1] //后面弹出POP_BACK
		}
		queue = append(queue,i) //后面插入PUSH_BACK
	}
	for i :=0; i< k; i++{
		push(i)
		fmt.Printf("%T,queue=%+v,i=%d,nums[i]=%d\n",queue, queue,i,nums[i])

	}
	fmt.Println("=============")
	len := len(nums)
	res := make([]int,1,len-k+1)
	res[0] = nums[queue[0]]

	for i:=k;i<len;i++{
		push(i)

		//fmt.Printf("%T,queue=%+v,i=%d,nums[i]=%d\n",queue, queue,i,nums[i])
		// i - k 这个 不太明白, i为len, k为窗口大小， i-k 为nums的长度减去窗口大小，为的是怕越界
		// queue[0]存储的是nums的索引，最大不超过 nums的长度减去窗口大小
		for queue[0] < i - k{ //滑动窖口,如果是递减的，每次得把前面最大弹出去去，不然队列会越来越长
			fmt.Printf("queue=%v,queue[0]=%d,nums[queue[0]]=%d,i=%d,k=%d,i-k=%d\n",queue,queue[0],nums[queue[0]], i,k,i-k)
			queue=queue[1:] //前面POP_FRONT
		}
		res= append(res,nums[queue[0]])
	}

	return res
}
func TestMSW() []int{
	return maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7,8,9,10,7,6,5,4,3,2,1}, 3)
	//return maxSlidingWindow([]int{1,-1}, 1)
}