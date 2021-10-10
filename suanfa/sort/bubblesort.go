package sort

import "fmt"

//从先到后逐步构建有序序列，对于未排序的数据，在已排序的序列中，从后向前扫描找到相应的位置
//第一层循环从1开始，默认第1个元素为有序的
//func insertionSort(arr []int){
//	len := len(arr)
//
//	for i :=1; i < len; i++{
//		for j:= 0; j< i; j++{
//			if arr[j] > arr[i] {
//				arr[j], arr[i] = arr[i],arr[j]
//			}
//		}
//	}
//}

//每次找最小值，然后放到数组的开始位置
//func selectionSort(arr []int)  {
//	len := len(arr)
//	for i:=0; i< len-1; i++ {
//		minIndex := i
//		for j:= i+1; j < len; j++ {
//			if arr[j] < arr[minIndex] {
//				arr[j], arr[minIndex] = arr[minIndex],arr[j]
//			}
//		}
//	}
//}
//双层嵌套循环，每次查看相邻的元素，如果逆序，则交换，一次会把最大的冒泡到最后
func bubbleSort(arr []int)  {
	len := len(arr)
	for i := 0; i< len -1; i++ {
		for j :=0; j < len-i-1; j++ {
			fmt.Printf("arr[%d] > arr[%d]\n", j,  j+1)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("++++++++++++++++++")

	}
}
func TestBubbleSort()  {
	n := 5
	for i:=n; i >=0; i--{
		for j:=i;j<n;j++{
			fmt.Printf("i=%d,j=%d\t",i,j)
		}
		fmt.Println()
	}
	//arr := []int{2,7,2,6,1,4,5}
	//bubbleSort(arr)
	//fmt.Println(arr)
}