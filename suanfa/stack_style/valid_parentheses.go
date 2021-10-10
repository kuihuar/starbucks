package stack_style

//
// 最近相关性-栈
// 先来后到-队列
// 暴力 不断的从头循环到尾，发现匹配的括号就替换成空
// 1. ((({{[]}})))
// 2. ((({{}})))
// 3. ((({})))
// 4. ((()))
// 5. (())
// 6. ()
// 7. 空-》合法
 // O(n^2)
//func isValid(s string) bool {
//	dic
//
//}

// 栈
// O(n) 扫一次
func IsValid1(s string) bool {
	parentheses := map[byte]byte{
		'{':'}',
		'[':']',
		'(':')',
	}
	stack := []byte{}
	for i:=0; i<len(s);i++{
		if _,ok := parentheses[s[i]];ok { //左括号
			stack = append(stack,parentheses[s[i]]) //如果是左括号，栈里放一个右括号
		}else {
			//栈为空的时间来了右括号K者 匹配不了
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
