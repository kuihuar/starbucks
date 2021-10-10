package datastructures

import "fmt"

type minHeapSort struct {
	arr []int
}
func NewMinHeapSort(arr []int) *minHeapSort {
	minheap := &minHeapSort{
		arr: arr,
	}
	return minheap
}
func (m *minHeapSort) leftChildIndex (index int)int{
	return 2 * index + 1
}
func (m *minHeapSort) rightChildIndex (index int)int{
	return 2 * index + 2
}
func (m *minHeapSort) buildMinHeapSort (size int){
	for index := size / 2 -1; index >=0; index-- {
		m.downHeapifySort(index, size)
	}
}
func (m *minHeapSort) leaf(index int, size int) bool {
	if index >= size / 2 && index <=size {
		return true
	}
	return false
}
func (m *minHeapSort) swap(first int, second int)  {
	m.arr[first] , m.arr[second] = m.arr[second],m.arr[first]
}
func (m *minHeapSort)downHeapifySort(current, size int)  {
	if m.leaf(current, size) {
		return
	}
	smallest := current
	leftChildIndex := m.leftChildIndex(current)
	rightChildIndex := m.rightChildIndex(current)
		 if leftChildIndex < size && m.arr[leftChildIndex]< m.arr[smallest] {
		 	smallest = leftChildIndex
		 }
		 if rightChildIndex < size && m.arr[rightChildIndex] < m.arr[smallest]{
		 	smallest = rightChildIndex
		 }
		 if smallest != current {
		 	m.swap(current, smallest)
		 	m.downHeapifySort(smallest, size)
		 }
}
func (m *minHeapSort)sort(size int)  {
	m.buildMinHeapSort(size)
	for i := size -1 ; i>0; i-- {
		m.swap(0,i)
		m.downHeapifySort(0,i)
	}
}
func (m *minHeapSort) print()  {
	for _,val := range m.arr {
		fmt.Println(val)
	}
}
func TestHeapSort()  {
	input := []int{6,5,3,7,2,8,1,15	}
	minheap := NewMinHeapSort(input)
	fmt.Printf("%+v",minheap.arr)
	minheap.sort(len(input))
	minheap.print()
}