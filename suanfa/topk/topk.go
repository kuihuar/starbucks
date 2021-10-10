package topk

import (
	"container/heap"
	"fmt"
)

//小项堆和哈希实现
type IHeap [][2]int

func (h IHeap)Len() int {
	return len(h)
}
func (h IHeap) Less (i, j int) bool{
	//小项堆 小于
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *IHeap) Push(x interface{})  {
	*h = append(*h,x.([2]int))
}
func (h *IHeap)Pop() interface{} {
	old := *h
	len := len(old)
	x := old[len-1]
	*h = old[0:len-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _,num := range nums{
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	//小顶堆，把前面频率小的POP出去，剩下的就是频率高的
	for key, value := range occurrences {
		heap.Push(h, [2]int{key,value})
		if h.Len() > k { //超过LEN，就把前面POP出去，POP出去的都是频率较小的
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//var res []int
	for i:=0; i<k;i++ {
		fmt.Println(k-i-1)
		res[k-i-1] = heap.Pop(h).([2]int)[0]//直接倒序排了
		//res[i] = heap.Pop(h).([2]int)[0]
		//res = append(res, heap.Pop(h).([2]int)[0])
		//res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}

func Test()  {
	res := topKFrequent([]int{1,1,1,1,2,2,3,55,55,55,55,55,55}, 2)
	fmt.Println(res)
}