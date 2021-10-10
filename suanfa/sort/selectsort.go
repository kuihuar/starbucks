package sort



//每次找最小值，然后放到数组的开始位置
func selectionSort(arr []int)  {
	len := len(arr)
	for i:=0; i< len-1; i++ {
		minIndex := i
		for j:= i+1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex],arr[j]
			}
		}
	}
}
