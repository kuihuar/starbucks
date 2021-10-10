package recircution

import "fmt"

func ThreeWayQuickSort()  {
	//data := []int{34,33,12,78,55,45,55,65,55,21,1,98,100}
	data := []int{34,33,12,78,21,1,98,100}
	sortTWQS(data,0, len(data)-1)
	fmt.Println(data)
}
func sortTWQS(data []int, low, hight int)  {
	if low >= hight{
		return
	}
	less, great  := partitionTWQS(data, low, hight)
	sortTWQS(data, low, less -1)
	sortTWQS(data, great +1, hight)
}
func partitionTWQS(data []int, low, hight int) (int, int) {
	pivot := hight
	less := low
	great:= hight
	i := low

	for i <= great {
		if data[i] < data[pivot] {
			data[i],data[less] = data[less],data[i]
			less++
			i++
		}else if data[i] > data[pivot] {
			data[i],data[great] = data[great], data[i]
			great--
		}else{
			i++
		}
	}
	return less, great
}