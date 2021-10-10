package search
//是否完全平方


func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)/2
		if mid * mid == num {
			return true
		}else if mid* mid > num {
			right = mid-1
		}else {
			left = mid+1
		}
	}
	return false
}