package stack_style

import "fmt"

// 两种方法
// 1. 用两个栈去实现， 一个保存正常操作，另gwh保存最小元素
// 1. 一个栈保存两个数据域，val 以及min value
type MinStack struct {
	val[]int
	min[]int

}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	this.val = append(this.val,val)
	//如果min栈为空，或者小于等于当前最小栈的数
	//不然插入相同的val时，会把当前最小栈越界
	if len(this.min) == 0  || val <= this.GetMin(){ // 这里必须是小于等于
		this.min = append(this.min,val)
	}
}


func (this *MinStack) Pop()  {
	item := this.val[len(this.val)-1]
	this.val = this.val[:len(this.val)-1]
	if this.GetMin() == item {
		this.min = this.min[:len(this.min)-1]
	}
}


func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func TestMinStack() int {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	return 1
}