?package heap

//1. sort  NlogN
//2. heap N logK
//3. quick-sort


// 大根堆，前K小
// 小根堆，前K大


// 小顶堆
import "container/heap"
type IntHeap []int
func(h IntHeap)Len() int {
	return len(h)
}
func(h IntHeap)Less(i,j int)bool{
	return h[i] > h[j]
}
func(h IntHeap)Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}
func (h *IntHeap)Push(x interface{}){
	*h = append(*h, x.(int))
}
func (h *IntHeap)Pop() interface{} {
	res := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return res
}
func getLeastNumbers(arr []int, k int) []int {
	h := &IntHeap{}
	for _,val :=range arr {
		h.Push(val)
	}
	heap.Init(h)
	var res []int
	for h.Len()>k{
		res = append(res, (h.Pop()).(int))
	}
	return res
}
