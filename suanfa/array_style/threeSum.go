package array_style

import (
	"sort"
)

//j暴力三重
//func ThreeSum(nums []int) [][]int {
//	len:= len(nums)
//	var res [][]int
//	for i:=0; i<len -2; i++{
//		for j:=i+1; j < len-1; j++ {
//			for k:= j+1; k<len;k++{
//				fmt.Printf("i=%d\t,j=%d\t, k=%d\n",i,j,k)
//				fmt.Printf("nums[i]=%d\t,nums[j]=%d\t, nums[k]=%d\n",nums[i],nums[j],nums[k])
//				if nums[i] + nums[j] - nums[k] == 0 {
//					res = append(res, []int{i,j,k})
//				}
//			}
//		}
//	}
//	return res
//}
// hash + loop
//func ThreeSum(nums []int) [][]int {
//	len:= len(nums)
//
//	var res [][]int
//	var exitsMap = make(map[string][]int)
//	for i:=0; i< len-2; i++{
//		target := -nums[i]
//		build := make(map[int]int)
//		for j:=i+1;j < len;j++{
//			componet := target - nums[j]
//			var keyStr string
//			if val,ok := build[componet]; ok {
//				item:= []int{nums[i],val,nums[j]}
//				sort.Ints(item)
//				for _, v := range item{
//					keyStr += strconv.Itoa(v) + " "
//				}
//				exitsMap[keyStr] = item
//			}
//			build[nums[j]] = nums[j]
//		}
//	}
//	for _,val := range exitsMap {
//		res = append(res, val)
//	}
//	return res
//}
//hash + 双指针
func ThreeSum(nums []int) [][]int {
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

func ThreeSum1(nums []int) [][]int {
	len:= len(nums)
	var res [][]int
	uniqueMap := make(map[[3]int]struct{})
	for i:=0; i<len -2; i++{
		for j:=i+1; j < len-1; j++ {
			for k:= j+1; k<len;k++{
				if nums[i] + nums[j] - nums[k] == 0 {
					sliceThreeNum := [3]int{nums[i] ,nums[j] , nums[k]}
					sort.Ints(sliceThreeNum[:])
					if _, ok := uniqueMap[sliceThreeNum];!ok{
						res = append(res, []int{i,j,k})
						uniqueMap[sliceThreeNum]= struct{}{}
					}
				}
			}
		}
	}
	return res
}
func ThreeSum2(nums []int) [][]int {
	len:= len(nums)
	var res [][]int
	sort.Ints(nums)
	for i:=0; i<len -2; i++{
		for j:=i+1; j < len-1; j++ {
			for k:= j+1; k<len;k++{
				if nums[i] + nums[j] - nums[k] == 0 {
					res = append(res, []int{nums[i] ,nums[j] , nums[k]})
				}
			}
		}
	}
	return res
}
//二分查找
func ThreeSum3(nums []int) [][]int {
	len:= len(nums)
	var res [][]int
	sort.Ints(nums)
	unique := make(map[[3]int]struct{})
	for i:=0; i<len; i++{
		target := -nums[i]
		left := i+1
		right := len-1 //right 指的是下标一定要减1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				item := [3]int{nums[i], nums[left], nums[right]}
				if _,ok := unique[item];!ok { //去重方法之一
					unique[item] =struct{}{}
					res = append(res,item[:])
				}
				left++
				right--
			}
			if sum < target {
				right--
			}else{
				left ++
			}
		}
	}
	return res
}

func ThreeSum4(nums []int) [][]int {
	len:= len(nums)
	var res [][]int
	sort.Ints(nums)
	var target int
	for i:=0; i<len; i++{
		left := i+1
			target = -nums[i]
		right := len-1 //right 指的是下标一定要减1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				res = append(res,[]int{nums[i], nums[left], nums[right]})
				//fmt.Println(res)
				left++
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				right--
				for left< right && nums[right] == nums[right-1]{
					right--
				}
			}else if sum < target {
				left++
			}else{
				right--
			}
		}
	}
	return res
}