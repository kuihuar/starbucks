package recircution

import "fmt"

//func generateParenthesis(n int) []string {
//	generate(0, 0,n,"")
//	return nil
//}
//var in int
//func generate(left int, right int, total int, currentParenthes string)  {
//	//terminator
//	if left == total && right == total {
//		fmt.Println(currentParenthes)
//		return
//	}
//	//process current logic: 加左括号或加右括号
//	s1 := currentParenthes + "("
//	//drill down
//	if left < total {
//		generate(left +1, right,total, s1)
//	}
//	s2 := currentParenthes + ")"
//	if left > right{
//		generate(left, right+1,total, s2)
//	}
//	in++
//	fmt.Println(in)
//	//reverse states
//	//s1和s2都是局部变量，不用reverse去清除
//}

func generateParenthesis(n int) []string {
	generate(0,2*n,"")
	return nil
}
var in int
func generate(level int,  total int, currentParenthes string)  {
	//fmt.Printf("level:%d\n",level)
	//terminator
	if level >= total {
		fmt.Println(currentParenthes)
		return // 为什么这个return 是回溯
	}
	//process current logic: 加左括号或加右括号
	s1 := currentParenthes + "("
	//drill down
	generate(level +1,total, s1)
	s2 := currentParenthes + ")"
	//fmt.Println("middle")
	generate(level +1, total, s2)
	//generate(level +1, total, s2)
	in++
	fmt.Println(in)
	//reverse states
	//s1和s2都是局部变量，不用reverse去清除
}

func TestGP() []string {

	return generateParenthesis(3)
}