package recircution

import "fmt"

func generateParenthes2(n int) []string {
	res := []string{}
	_generate(0,0,n, "", &res)
	return res

}
func _generate (left int, right int, n int, path string, res *[]string) {
	if left ==n && right == n{
		*res = append(*res,path)
		return
	}
	if left < n{
		_generate(left+1, right, n, path+"(", res)
	}
	if right <left {
		_generate(left, right+1, n, path+")", res)
	}
}
func 	TestGPS2()  {
	res := generateParenthes2(3)
	fmt.Println(res)
}