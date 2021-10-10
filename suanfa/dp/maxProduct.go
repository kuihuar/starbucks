package dp

import "fmt"

func maxProduct(nums []int) int {

	max,min,res := nums[0],nums[0],nums[0]
	for i:=1; i< len(nums); i++{
		if nums[i] <0 {
			max,min= min,max
		}
		if max * nums[i] > nums[i] {
			max = max * nums[i]
		}else{
			max = nums[i]
		}
		if min * nums[i] < nums[i] {
			min = min * nums[i]
		}else{
			min = nums[i]
		}
		if max > res {
			res = max
		}
	}
	return res
}

func TestMaxProduct()  {
	p := []int{2,3,-2,4}
	res := maxProduct(p)
	fmt.Println(res)
}