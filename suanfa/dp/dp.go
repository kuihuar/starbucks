package dp

import (
	"fmt"
)

func fib(n int) int {
	if n<=1 {
		return n
	}else{
		return fib(n-1) + fib(n-2)
	}
}

// 从指数降到O(n)
func fib1(n int) int {
	var helper func(int) int

	mem := make(map[int]int)
	helper = func(i int) int {
		if i <= 1 {
			return i
		}
		if mem[i] == 0 {
			mem[i] = helper(i-1) + helper(i-2)
		}
		return mem[i]
	}
	res := helper(n)
	return res
}
func fibHelper(n int, mem *map[int]int) int {
	if n<=1 { //递归终止条件
		return n
	}
	if (*mem)[n] == 0{
		fmt.Println("n= ",n)
		(*mem)[n] = fib1(n-1) + fib1(n-2)
	}
	return (*mem)[n]
}
//循环处理，n得加2，不然越界
func fib2(n int) int {
	s := make([]int,n+2,n+2)
	s[0]=0
	s[1]=1
	for i:=2;i<=n;i++{
		s[i] = s[i-1]+ s[i-2]
	}
	return s[n]
}
func TestDp()  {
	res :=  fib1(5)
	fmt.Println(res)
}