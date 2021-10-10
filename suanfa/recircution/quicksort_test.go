package recircution

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{46,9,34,11,33,13,22,77,45,8,10,20}
	ans := QuickSort(data)
	//fmt.Println(ans)
	//libSort.Slice(data, func(i, j int) bool {
	//	return data[i] < data[j]
	//})
	fmt.Println(ans)
}