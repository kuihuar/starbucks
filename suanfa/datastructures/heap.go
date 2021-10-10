package datastructures

import "fmt"

type minheap struct{
	heapArray []int
	size int
	maxsize int
}

func newMinHeap(maxsize int)*minheap {
	minheap := &minheap{
		heapArray: []int{},
		size:     0,
		maxsize:  maxsize,
	}
	return minheap
}
func (m *minheap)leaf(index int) bool {
	if index >= m.size / 2 && index <=m.maxsize {
		return true
	}
	return false
}
func (m *minheap)parent(index int) int {
	return (index - 1 ) / 2
}

func (m *minheap) leftchild(index int) int {
	return 2 * index + 1
}
func (m *minheap) rightchild(index int) int {
	return 2 * index +2
}
func (m *minheap) insert(item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("heap is full")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}
func (m *minheap)swap(first, second int)  {
	m.heapArray[first], m.heapArray[second] = m.heapArray[second], m.heapArray[first]
}
//向上堆化
func (m *minheap) upHeapify(index int){
	for m.heapArray[index] < m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
	}
}
func (m *minheap)buidMinheap()  {
	for index := (m.size / 2) - 1; index >=0; index-- {
		m.downHeapify(index)
	}
}
func (m *minheap)buidMinheapWithsize(size int)  {
	for index := (size / 2) - 1; index >=0; index-- {
		m.downHeapify(index)
	}
}
//最下面有叶子节点的
func (m *minheap)downHeapify(current int)  {
	if m.leaf(current){
		return
	}
	smallest := current
	leftChildIndex := m.leftchild(current)
	rightChildIndex := m.rightchild(current)
	if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest]{
		smallest = leftChildIndex
	}
	if rightChildIndex < m.size && m.heapArray[rightChildIndex] < m.heapArray[smallest] {
		smallest = rightChildIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
	return
}
func (m *minheap)remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size -1]
	m.heapArray = m.heapArray[:m.size-1]
	m.size--
	m.downHeapify(0)
	return top
}
func TestMinheap(){
	inputArray := []int{6,5,3,7,2,8, -1}
	minHeap := newMinHeap(len(inputArray))
	for i:=0; i<len(inputArray); i++{
		minHeap.insert(inputArray[i])
	}
	minHeap.buidMinheap()
	fmt.Println(minHeap.heapArray)
	for i:=0; i < len(inputArray); i++{
		fmt.Println(minHeap.remove())
	}
	fmt.Scanln()
}
func (m *minheap) print() {
	for _ ,val := range m.heapArray {
		fmt.Println(val)
	}
}