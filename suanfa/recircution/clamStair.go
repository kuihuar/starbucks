package recircution

import "fmt"

//f(n) = f(n-1) + f(n-2)
//func ClimbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	return ClimbStairs(n-1) + ClimbStairs(n-2)
//}

// 保存最近的的三个数
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i < n+1; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}


//傻递归
func ClimbStaris(i int, total int)  int {
	if i > total {
		return 0
	}
	if i == total {
		return 1
	}
	return ClimbStaris(i +1, total) + ClimbStaris(i +2, total)
}

//记忆化递归
func climbStarisWithMemory(total int) int{
	mem := make([]int,total+1,total+1)
	return climbStarisWithMemoryHelper(0, total, mem)
}
func climbStarisWithMemoryHelper(i,total int, mem []int)  int {
	if i > total {
		return 0
	}
	if i == total {
		return 1
	}
	if mem[i] > 0 {
		return mem[i]
	}
	mem[i] = climbStarisWithMemoryHelper(i +1,total, mem) + climbStarisWithMemoryHelper(i +2, total,mem)
	return mem[i]
}

//动态规划

func ClimbStarisDp(totalStairs int) int {
	if totalStairs == 1 {
		return 1
	}
	var dp = make([]int, totalStairs +1)
	dp[1]= 1
	dp[2]= 2
	fmt.Println(dp)
	for i :=3;i<= totalStairs;i++{
		fmt.Println(i)
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println(dp)
	return dp[totalStairs]
}
func Test() int {
	res:=  ClimbStarisDp(4)
	//res:=  ClimbStairs(4)
	return res
}