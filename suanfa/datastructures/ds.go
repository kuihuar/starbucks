package datastructures

import (
	"container/list"
	"fmt"
)
type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue()(value string)  {
	c.queue.PushBack(value)
	return
}
func (c *customQueue) Dequeue() error  {
	if c.queue.Len() >0 {
		//if val, ok := c.queue.Front().Value.(string); ok {
		//	return val,nil
		//}
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return  fmt.Errorf("not found")

}
func(c *customQueue) Front()(string, error){
	if c.queue.Len() >0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val,nil
		}

	}
	return  "" ,fmt.Errorf("not found")
}
func (c *customQueue) Size()int{
	return c.queue.Len()
}
func(c *customQueue)Empty()bool{
	return c.queue.Len()==0
}