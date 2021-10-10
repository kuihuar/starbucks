package search

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	if x == 0 || x==1{
		return x
	}
	left, right:= 1,x
	for left <= right{
		mid := left + (right-left) / 2
		if mid * mid > x {
			right = mid -1
		}else{
			left = mid+1
		}
	}
	return right
}

// 牛顿迭代法
func mySqrt1(x int) int {
	if x == 0{
		return 0
	}
	c,x0 := float64(x),float64(x)
	for{
		xi:= 0.5 * (x0 + c/x0)

		//cur = (cur +x / cur) / 2
		if math.Abs(x0 -xi) < 1e-6 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
func TestSqrt()  {
	res := mySqrt3(5)
	fmt.Println(res)
}
//stefanPochmann
func mySqrt2(x int) int {
	r := float64(x)
	xf := float64(x)
	for r * r > xf {
		r = (r + xf /r) / 2
	}
	return int(r)
}
func mySqrt3(x int) int {
	r:=x
	for r * r > x {
		r = (r + x /r) / 2
	}
	return r
}