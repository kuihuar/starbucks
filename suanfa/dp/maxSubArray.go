package dp
func maxSubArray(nums []int) int {
	for i:=1; i< len(nums);i++{
		if nums[i] <  nums[i] + nums[i-1]{
			nums[i] = nums[i] + nums[i-1]
		}
		//if nums[i-1] > 0 {
		//	nums[i] = nums[i] + nums[i-1]
		//}
		//dp[i] = nums[i] + max(0,dp[i-1])
		//dp[i]= max(nums[i]+0, nums[i]+nums[i-1])
 	}
 	//return max(nums)
 	return 0
}

func maxSubArray1(nums []int) int {
	maxValue := nums[0]
	for i:=0; i< len(nums);i++{
		if nums[i] <  nums[i] + nums[i-1]{
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > maxValue{
			maxValue = nums[i]
		}
	}
	return maxValue
}
//洋释
func maxSubArray4(nums []int) int {
	var dp []int
	copy(dp,nums)

	for i:=0; i< len(nums);i++{
		// dp[i-1] 之前的
		// nums[i] 当前元素
		// max(0 + dp[i-1])  包含之前
		//dp[i] = max(0 + dp[i-1]) + nums[i]
	}
	//return max(dp)
	return 0
}
func maxSubArray3(nums []int) int {
	var dp []int
	copy(dp,nums)
	//dp[i]=max(nums[i], nums[i] + dp[i-1])
	//最大子序和=当前元素自身最大，或者包含之前以后最大
	for i:=0; i< len(nums);i++{
		if nums[i] <  nums[i] + nums[i-1]{
			dp[i] = nums[i] + nums[i-1]
		}
	}
	var maxValue int
	for i:=0; i< len(nums);i++{
		if nums[i] > maxValue{
			maxValue = nums[i]
		}
	}
	return maxValue
}