package datastructures

import (
	"fmt"
)

type maxheap struct {
	heapArray []int
	size int
	maxsize int
}
func newMaxheap(maxsize int ) *maxheap {
	return &maxheap{
		heapArray: []int{},
		size:      0,
		maxsize:   maxsize,
	}
}

func (m *maxheap) leaf (index int) bool{
	if index >= m.size /2 && index < m.size {
		return true
	}
	return false
}
func (m *maxheap) parent(index int) int {
	return (index-1) / 2
}
func(m *maxheap) leftChild(index int) int {
	return 2 * index +1
}
func (m *maxheap)rightChild(index int) int {
	return 2 * index + 2
}
func(m *maxheap)kthChild(index,k int)int {
	return 2 * index +k
}
func (m *maxheap) swap(first, second int){
	m.heapArray[first], m.heapArray[second] = m.heapArray[second], m.heapArray[first]
}

func (m *maxheap) insert (item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("heap is ful")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}
func (m *maxheap)upHeapify(index int) {
	for m.heapArray[index] > m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}
func (m *maxheap)upHeapify1(index int) {
	insertValue := m.heapArray[index]
	//index > 0 表示index没能触及到根
	for index > 0  && insertValue > m.heapArray[m.parent(index)]{
		m.heapArray[index] = m.heapArray[m.parent(index)]
		index = m.parent(index)
	}
	m.heapArray[index] = insertValue
}
func (m *maxheap)downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)
	if leftChildIndex < m.size && m.heapArray[leftChildIndex] > m.heapArray[largest] {
		largest = leftChildIndex
	}
	if rightChildIndex < m.size  && m.heapArray[rightChildIndex] > m.heapArray[largest] {
		largest = rightChildIndex
	}
	if largest!= current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}
func (m *maxheap) buildMaxheap()  {
	for index := m.size / 2 -1; index >=0; index-- {
		m.downHeapify(index)
	}
}
func (m *maxheap)delete(x int) int {
	element := m.heapArray[x]
	m.heapArray[x] = m.heapArray[m.size-1] //堆尾元素放到堆顶
	m.size--
	m.downHeapify(x)
	return element
}
func (m *maxheap) downHeapify1(index int) {
	var child int
	temp := m.heapArray[index]
	for m.kthChild(index,1) < m.size { //孩子节点不能超出堆的大小
		child = m.maxChild(index)
		if temp >= m.heapArray[child]{ //找到以后就退出for, 执行 m.heapArray[index] = temp
			break
		}
		m.heapArray[index] = m.heapArray[child] //儿子的值赋值给自己
		index = child //否则去更大的Child走
	}
	m.heapArray[index] = temp //堆尾放远来的元素赋值到相应在的位置去
}
func (m *maxheap)maxChild(index int) int{
	leftChild := m.kthChild(index,1)
	rightChild := m.kthChild(index,2)
	if m.heapArray[leftChild] > m.heapArray[rightChild] {
		return leftChild
	}else{
		return rightChild
	}
}
func (m *maxheap)remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:m.size-1]
	m.size--
	m.downHeapify(0)
	return top
}

func TestMaxheap() {
	inputArray := []int{6,5,3,7,2,8}
	maxHeap := newMaxheap(len(inputArray))
	for i:=0; i< len(inputArray); i++ {
		maxHeap.insert(inputArray[i])
	}
	maxHeap.buildMaxheap()
	fmt.Println("The Max heap is:")
	for i:=0; i<len(inputArray); i++ {
		fmt.Println(maxHeap.remove())
	}
}
