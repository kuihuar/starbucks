package suanfa
//有序数组
// 1 二分查找
func twoSum(numbers []int, target int) []int {
	res := make([]int,2)
	for i:=0; i< len(numbers); i++{
		searchNum:= numbers[i]
		index := binnarySearch(numbers, i+1, len(numbers)-1, target - searchNum)
		if index == -1{
			res[0] = index+1
			res[1] = i+1
		}
	}
	return res
}
func binnarySearch(numbers []int, lower, upper, target int) int{
	for lower <= upper {
		midder := (lower + upper) / 2
		if numbers[midder] == target {
			return midder
		}
		if numbers[midder] < target {
			lower = midder +1
		}else{
			upper = midder -1
		}
	}
	return -1
}
// 2 双指针 O(n)
func twoSum1(numbers []int, target int) []int {
	left :=0
	right := len(numbers) -1
	res := make([]int, 2,2)
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			res[0] = left +1
			res[0] = right+1
		}
		if sum < target {
			left++
		}else{
			right--
		}
	}
	return res
}