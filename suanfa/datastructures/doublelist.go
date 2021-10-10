package datastructures

import "fmt"

//doublelist 相当于 deque 双端队列
// addFirst removeFirst getFirst
// addLast removeLast getLast

type node struct {
	data string
	prev *node
	next *node
}
type doubleList struct {
	len int
	tail *node
	head *node
}
func initdoublelist() *doubleList {
	return &doubleList{}
}

func(d *doubleList)AddFront(data string){
	newNode := &node{data: data}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	}else{
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}
func (d *doubleList)AddBack(data string)  {
	newNode := &node{data: data}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	}else{
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
		d.len++
		return
	}
}
func (d *doubleList)TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("empty")
	}
	curr := d.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	return  nil
}
func (d *doubleList)TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("empty")
	}

	curr := d.tail
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.prev
	}
	return nil
}
func (d *doubleList)Size() int {
	return d.len
}
func (d *doubleList)Reverse()  {
	current := d.head
	var nextInList *node
	d.head, d.tail = d.tail, d.head // 交换列的的prev和next
	for current != nil { //交换当前节点的prev和next
		nextInList = current.next
		current.next, current.prev = current.prev, current.next
		current = nextInList
	}
}