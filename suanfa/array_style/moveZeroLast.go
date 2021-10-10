package array_style

//数组里，把0移动到数组末尾
//给定一个数组，一个函数将0移动到末尾，保持非零元素的相对顺序

//
func MoveZeroes(nums []int)  {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j{
				nums[i] = 0
			}
 			j++
		}
	}
}
func MoveZeroes1(nums []int)  {
	j := 0 //(非零元素的位置)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] , nums[i] = nums[i], nums[j]
			j++
		}
	}

}

func MoveZeros2(nums []int) {
	count := 0
	for i:=0; i<len(nums); i++{
		if nums[i] == 0 {
			count++
		}else{
			nums[i-count] = nums[i]
		}
	}
	for i:=len(nums)- count;i<len(nums);i++ {
		nums[i]=0
	}
}