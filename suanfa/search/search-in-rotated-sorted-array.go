package search

import "fmt"

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		//当[0,mid]有序时向右半部分规约
		if nums[0] <= nums[mid] && (target > nums[mid] || target < nums[0]){
			left = mid +1
			//当[0,mid]发生旋转时向右规约
		}else if target > nums[mid] && target < nums[0]{
			left = mid +1
		}else {//向左
			right = mid
		}
	}
	if left == right && nums[left] == target {
		return left
	}
	return -1
}
func Testsearch(){
	res := search([]int{4,5,6,7,0,1,2},0)
	fmt.Println(res)
}