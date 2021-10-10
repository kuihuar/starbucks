package datastructures

import "container/heap"

//import "container/heap"
//
//type IntHeap []int
//func(h IntHeap)Len() int {
//	return len(h)
//}
//func(h IntHeap)Less(i,j int)bool{
//	return h[i] > h[j]
//}
//func(h IntHeap)Swap(i,j int){
//	h[i],h[j]=h[j],h[i]
//}
//func (h *IntHeap)Push(x interface{}){
//	*h = append(*h, x.(int))
//}
//func (h *IntHeap)Pop() interface{} {
//	res := (*h)[len(*h)-1]
//	(*h) = (*h)[:len(*h)-1]
//	return res
//}
//func getLeastNumbers(arr []int, k int) []int {
//	h := &IntHeap{}
//	for _,val :=range arr {
//		h.Push(val)
//	}
//	heap.Init(h)
//	var res []int
//	for h.Len()>k{
//		res = append(res, (h.Pop()).(int))
//	}
//	return res
//}
func getLeastNumbers(arr []int, k int) []int {
	h:=&heapInt{}
	heap.Init(h)
	for _,v:=range arr {
		if h.Len()<k {
			heap.Push(h,v)
		}else {
			if h.Peek()>v {
				heap.Pop(h)
				heap.Push(h,v)
			}
		}
	}
	return *h
}
type heapInt []int

//Less  小于就是小跟堆，大于号就是大根堆
func (h *heapInt)Less(i,j int)bool {return (*h)[i]>(*h)[j]}
func (h *heapInt)Swap(i,j int) {(*h)[i],(*h)[j]=(*h)[j],(*h)[i]}
func (h *heapInt)Len() int {return len(*h)}
func (h *heapInt)Push(x interface{}) {
	*h=append(*h,x.(int))
}
func (h *heapInt)Pop() interface{} {
	t:=(*h)[len(*h)-1]
	*h=(*h)[:len(*h)-1]
	return t
}
func (h *heapInt)Peek() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return 0
}
func TestLeastNumber()  []int{
	//arr := []int{4,5,1,6,2,3,7,8}
	//arr := []int{3,2,1}
	arr := []int{0,0,0,2,0,5}
	return getLeastNumbers(arr, 0)
}