package datastructures

import (
	"container/list"
	"fmt"
)

//last in firt out
type customStack struct {
	stack *list.List
}

func(c *customStack)Push(value string) {
	c.stack.PushFront(value)
}
func(c *customStack)Pop()error {
	if c.stack.Len() >0 {
		ele:=c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("not found")
}
func(c *customStack) Front()(string, error){
	if c.stack.Len() > 0{
		if val,ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
	}
	return "", fmt.Errorf("not found")
}
// 相当于Front
func (c *customStack)Peek() (string, error) {
	if c.stack.Len() > 0{
		if val,ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
	}
	return "", fmt.Errorf("not found")
}
func (c *customSet)Search() bool {
	return false
}