package heap

import (
	"container/heap"
	"fmt"
)

//大顶堆实现,使用优先队列实现堆

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func maxSlidingWindow(nums []int, k int)  []int{
	pq := make(PriorityQueue,k,k)
	for i:=0; i<k;i++{
		pq[i] = &Item{
			value:    nums[i],
			priority: nums[i],
			index:    i,
		}
	}
	heap.Init(&pq)
	fmt.Printf("%T, %+v\n",pq,pq[0])
	fmt.Printf("%T, %+v\n",pq,pq[1])
	fmt.Printf("%T, %+v\n",pq,pq[2])
	var res []int
	//fmt.Printf("%T, %+v\n",item,item)
	//item = heap.Pop(&pq).(*Item).value
	//fmt.Printf("%T, %+v\n",item,item)
	res =append(res,pq[0].value)
	fmt.Println(res)
	for i:=k; i<len(nums);i++{
		heap.Push(&pq, &Item{
			value:    nums[i],
			priority: nums[i],
		})
		fmt.Printf("%T, %+v\n",pq,pq[0])
		//这个是啥意思啊。不明白
		for pq[0].value < i -k{ //滑动窖口,如果是递减的，每次得把前面最大弹出去去，不然队列会越来越长
			heap.Pop(&pq)
			fmt.Println(pq[0])
		}
		res =append(res,pq[0].value)
	}
	return res
	// Some items and their priorities.
	//items := map[string]int{
	//	"banana": 3, "apple": 2, "pear": 4,
	//}
	//
	//// Create a priority queue, put the items in it, and
	//// establish the priority queue (heap) invariants.
	//pq := make(PriorityQueue, len(items))
	//i := 0
	//for value, priority := range items {
	//	pq[i] = &Item{
	//		value:    value,
	//		priority: priority,
	//		index:    i,
	//	}
	//	i++
	//}
	//heap.Init(&pq)
	//
	//// Insert a new item and then modify its priority.
	//item := &Item{
	//	value:    "orange",
	//	priority: 1,
	//}
	//heap.Push(&pq, item)
	//pq.update(item, item.value, 5)
	//
	//// Take the items out; they arrive in decreasing priority order.
	//for pq.Len() > 0 {
	//	item := heap.Pop(&pq).(*Item)
	//	fmt.Printf("%.2d:%s ", item.priority, item.value)
	//}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple

}

func TestMaxSlidewidow() error {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k:= 3
	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
	//maxSlidingWindow()
	return nil
}