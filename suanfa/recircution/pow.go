package recircution


//暴力法
//func MyPow(x float64, n int) float64 {
//	if n < 0 {
//		x = 1 / x
//		n = -n
//	}
//	var res float64 = 1
//	for i:=0; i< n; i++ {
//		res = res * x
//	}
//	return res
//}
//分冶法
func MyPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		return 1 /  myPowHelper(x , -n )
	}
	return myPowHelper(x,n)
}

func myPowHelper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	half := myPowHelper(x, n / 2)
	if n % 2 != 0 {
		return half * half
	}else{
		return x * half *half
	}
}