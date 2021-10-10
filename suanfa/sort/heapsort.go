package sort

//findMax deleteMax insert create
//可以迅速找到一堆数中的最大者或最小者
//find O(1)
//delete O(logN)
//insert/create O(logN) OR O(1)

//堆的不同实现，链表，数组，二叉树 （binary， leftist,binaomial, fibonacci）
//实现：
// 1。 完全二叉树
// 2。 树中的任意节点的值总是（大顶堆）大于子节点的值
// 2。 树中的任意节点的值总是（小顶堆）小于子节点的值

//因为是完全二叉树，所以用数组实现
//根无素索引为0
//索引为i的左孩子的索引为2*i+1
//索相为i的右孩子的索引为2*i+1
//索引为i的父结点的索引为 (i-1)/2
//根节点 heap[0]

//查找，直接返回heap[0]
//插入，先插和尾部，再重新维护堆 heapifyup ,时间复杂度最大就是树的深度 O(log2N)，上浮操作
//delete max 删除堆项， 将堆尾元素替换为该元素，并向下heapifyDown
// heapifyDown， 判断左右孩子的，取大的交换，一直到叶子
// 工程中，可以用priority_queue

//golang goog
import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func TestHeap() error {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	return nil
}
func Testfoobar() error {
	demo := make([]int,4,4)
	demo[0] = 5
	demo[1] = 8

	length := len(demo)
	//demo[length+1] =9
	demo[2] = 4
	fmt.Println(demo[2])
	fmt.Printf("%T,%+v,%d,%d", demo,demo,length,cap(demo))
	return nil
}

func heapsort(arr []int)  {
	h:= &IntHeap{}
	//h := &IntHeap{2, 1, 5}
	heap.Init(h)
	len := len(arr)
	for i:=0; i< len; i++{
		heap.Push(h,arr[i])
	}
	//heap.Push(h, 3)
	//fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	for i:=0; i< len;i++{
		arr[i] = heap.Pop(h).(int)
	}

}