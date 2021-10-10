package datastructures

import (
	"fmt"
)

type ele struct {
	name string
	next *ele
}
type singleList struct {
	len int
	head *ele
}
func initList()*singleList {
	return &singleList{}
}
func (s *singleList)AddFront(name string)  {
	ele := &ele{name: name}
	if s.head == nil {
		s.head = ele
	}else{
		ele.next = s.head
		s.head = ele
	}
	s.len++
	return
}
func (s *singleList)AddBack(name string)  {
	ele := &ele{name: name}
	if s.head == nil {
		s.head = ele
	}else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next =ele
	}
	s.len++
	return
}
func(s *singleList)RemoveFront()error{
	if s.head == nil {
		return fmt.Errorf("empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}
func(s *singleList)RemoveBack()error{
	if s.head == nil {
		return fmt.Errorf("empty")
	}
	var prev *ele
	current := s.head
	for current.next !=nil {
		prev = current
		current = current.next
	}
	prev.next = nil

	return  nil
}
func (s *singleList)Len() int {
	return s.len
}
func(s *singleList)Front()string {
	return s.head.name
}
func(s *singleList)Traverse()error{
	if s.head == nil {
		return fmt.Errorf("empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

func (s *singleList) ToArray() []string {
	var res []string
	current := s.head
	for current!= nil {
		res = append(res, current.name)
		current = current.next
	}
	return res
}
func TestSingleList(){
	singleList := initList()
	singleList.AddFront("D")
	singleList.AddFront("C")
	singleList.AddFront("B")
	singleList.AddFront("A")
	fmt.Println(singleList.Len())

	singleList.Traverse()
	arr := singleList.ToArray()
	fmt.Println(arr)
}
func(s *singleList) ToCircular (){
	current := s.head
	for current.next != nil {
		current = current.next
	}
	current.next = s.head
}
func (s *singleList) IsCircular() bool{
	if s.head == nil {//空的，没有数据的
		return true
	}
	//if s.head.next == s.head {//和上边的判断空的有什么不一样 ，s.head = ele 不一样的。这里如果有，就不是nil
	//
	//}
	//
	current := s.head.next
	for current != nil && current != s.head {
		current = current.next
	}
	return current == s.head
}
func (s *singleList)Removekth(k int) error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	if k > s.len {
		return fmt.Errorf("given num is greater list")
	}
	if k == 1 {
		s.head = s.head.next
	}else{
		var prev *ele
		current := s.head
		for i := 1; i<k;i++ {
			prev = current
			current = current.next
		}
		prev.next = current.next
	}
	s.len--
	return nil
}
func (s *singleList) removeKthFromEnd(k int) error {
	if s.head == nil {

	}
	if s.len < k{

	}

	if k == s.len {
		s.head = s.head.next
	}else{
		var prev *ele
		current := s.head

		for i:=1; i<s.len-k+1; i++ {
			prev = current
			current = current.next
		}
		prev.next = current.next
	}
	s.len--
	return nil
}