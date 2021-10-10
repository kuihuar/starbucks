package recircution

import "fmt"

func generateParenthes(n int) []string {
	res := []string{}
	var _generate func(left int, right int, toal int, path string)
	_generate = func(left int, right int, toal int, path string) {
		if left ==n && right == n{
			res = append(res,path)
			return
		}
		if left < n{
			_generate(left+1, right, n, path+"(")
		}
		if right <left {
			_generate(left, right+1, n, path+")")
		}
	}
	//fmt.Println(res)
	_generate(0,0,n, "")
	return res

}

func TestGPS()  {
	res := generateParenthes(3)
	fmt.Println(res)
}