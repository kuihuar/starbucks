//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：nums = [0]
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 3000 
// -10⁵ <= nums[i] <= 10⁵ 
// 
// Related Topics 数组 双指针 排序 👍 3675 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	len:= len(nums)
	var res [][]int

	for i:=0; i< len-2; i++{
		l, r := i+1, len-1//左指针//右指针
		if i >0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		for l < r {//两边夹逼
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i] , nums[l] , nums[r]})
				for l<r && nums[l] == nums[l+1] { //移动重复的
					l++
				}
				for l < r && nums[r] == nums[r-1] {//移动重复的
					r--
				}
				l++ //至少移动一次
				r-- //至少移动一次
			}else if sum < 0{
				l++
			}else{
				r--
			}
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
