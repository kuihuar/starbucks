package greedy
func canJump(nums []int) bool {
	endReachIndex := len(nums) -1

	for i:= len(nums) -1; i >=0; i-- {
		if nums[i] +i >=endReachIndex {
			endReachIndex = i
		}
	}
	return endReachIndex ==0
}
