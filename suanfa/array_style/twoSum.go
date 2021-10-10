package array_style

//暴力
//func TwoSum(nums []int, target int) []int {
//	len := len(nums)
//	//res := []int{}
//	var res []int
//	for i:=0; i<len -1; i++{
//		for j:=i+1; j < len; j++ {
//			if nums[i] + nums[j] == target{
//				res=append(res,i)
//				res=append(res,j)
//			}
//		}
//	}
//	return res
//}


//暴力，两个循环
func twoSum(nums []int, target int) []int {
	res := make([]int,0)
	for i:=0; i<len(nums) -2; i++{
		x := nums[i]//线性查找 O（n）
		for j :=i +1; j< len(nums)-1; j++ {
			if nums[j] == target -x {
				return []int{i,j}
			}
		}
	}
	return res
}


// hash
func TwoSum(nums []int, target int) []int {
	builds := make(map[int]int)
	for i:=0; i< len(nums); i++{
		complement := target - nums[i]

		if j,ok := builds[complement];ok{
			return []int{j,i}
		}
		builds[nums[i]] = i
	}
	return []int{0,0}
}