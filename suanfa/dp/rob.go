package dp

import "fmt"
//打家劫舍
func rob(nums []int) int {

	var max func(int, int)int
	max = func(i int, i2 int) int {
		if i> i2{
			return i
		}
		return i2
	}
	len := len(nums)

	dp := make([][]int, len)
	for i:=0; i< len;i++{
		dp[i] = make([]int,2)
	}

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < len; i++{
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0]+ nums[i]
	}
	return max(dp[len-1][0], dp[len-1][1])
}

func rob1(nums []int) int {
	var max func(int, int)int
	max = func(i int, i2 int) int {
		if i> i2{
			return i
		}
		return i2
	}
	len := len(nums)
	dp := make([]int,len)
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])
	for i:=2;i<len;i++{
		dp[i] = max(dp[i-1], nums[i] + dp[i-2])
	}
	return dp[len-1]
}
func rob2(nums []int) int {
	var max func(int, int)int
	max = func(i int, i2 int) int {
		if i> i2{
			return i
		}
		return i2
	}
	pre:=0
	now:=0
	for _,i := range nums{
		// dp[i] = max(dp[i-1], nums[i] + dp[i-2])
		pre,now = now,max(pre+i,now)
	}
	return now
}
func TestRob()  {
	//p := []int{1,2,3,1}
	p := []int{2,7,9,3,1}
	res := rob2(p)
	fmt.Println(res)
}