package devideconquer


//1. 循环n次，累乘，O(n)
//2. 分冶法 log(N)

func myPow(x float64, n int) float64 {

	// 1. teminator
	// 2. process split your big problem
	// 3. drill down subproblems, merge(subrest)
	// 4. reverse states if needed
	if n < 0 {
		x = 1 / x
		n = - n
	}
	return helper(x, n)
}
func helper(x float64, n int) float64 {

	if n == 0 { // teminator
		return x
	}
	// drill down subproblems,转化为子问题
	// 这里走了两步，1.拆分; 2. 转化为子问题


	//// x^10->2^10: x^5 -> 2*(2^2)
	//// problem: pow(x, n)
	//// subproblem:  subresult=pow(x, n/2)
	//// mrege
	//// if n%2 == 1{
	////		//odd
	////   result = subresult * subresult * x
	////  }else{
	////   // even
	////  result = subresult * subresult
	////}
	//当层处理的逻辑就是 n /2 ,直接drill down
	subresult := helper(x, n / 2) //转化成子问题，调用函数处理子问题


	//merge
	// 得到子问题之后，就把子问题合并成一个当前总的问题的一个结果
	//如何合并？
	if n % 2 ==0 {
		return subresult * subresult
	}else{
		return subresult * subresult * x
	}
}