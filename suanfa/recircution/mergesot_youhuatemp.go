package recircution

import (
	"fmt"
)

// 递推公式，sort(left,mid),sort(mid+1 right
//终止条件 left==right
func mergeSortYH(data []int)  {
	temp := make([]int,len(data) -1)
	sortYH(data,0, len(data) -1, temp)

}

func sortYH(arr []int, left, right int, temp []int)  {
	if left == right {
		return
	}
	mid :=  (left + right) / 2
	sortYH(arr, left, mid,temp)
	sortYH(arr, mid +1, right,temp)
	mergeYH2(arr, left, mid, right, temp)
}
func mergeYH(arr []int, left, mid, right int, temp []int)  {

	i:=left
	j:= mid+1
	tempPos:=left
	//左边数组和右边数组中的元素按顺序放到临时数组中 左边和右边都有无素
	for i<=mid && j<=right {
		if arr[i]<= arr[j] {
			temp[tempPos] = arr[i]
			tempPos++
			i++
		}else{
			temp[tempPos] = arr[j]
			tempPos++
			j++
		}
	}
	//剩余的数组（左边还有元素）
	for i <= mid {
		temp[tempPos] = arr[i]
		tempPos++
		i++
	}
	//剩余的数组（右边还有元素）
	for j <= right {
		temp[tempPos] = arr[j]
		tempPos++
		j++
	}
	fmt.Printf("arr=%v",arr)
	fmt.Printf("temp=%v\n",temp)
	//将临时数组拷贝给原数组
	//num := right-left + 1 区间的元素个数，有多少拷多少
	//num := right-left + 1
	for tempPos= left; tempPos< right;tempPos++{
		arr[left] = temp[tempPos]
		left++
	}
}
//https://ke.qq.com/webcourse/index.html#cid=3171403&term_id=103298623&taid=11370045451035723&vid=5285890809376429535
//将原始数组拷贝到临时数组，针对临时数组排序，归到到原始数组区间
func mergeYH2(arr []int, left, mid, right int, temp []int)  {

	for i:=left; i<=right;i++{
		temp[i] = arr[i] //拷贝到临时数组里
	}

	i:= left
	j:= right

	for k:=left ; k<=right;k++{
		if i == mid +1 { //左边没有元素,已经走到头了，将右边剩下的元素都拷贝过去
			arr[k] = temp[j]
			j++
		}else if j== right +1 {//右边索引已经走到头了，将左边剩下的元素都拷贝过去
			arr[k] = temp[i]
			i++
		}else if temp[i] <= temp[j] { //左边小于右边，两边都有元素
			arr[k] = temp[i]
			i++
		}else{//左边大于右边，两边都有元素
			arr[k] = temp[j]
			j++
		}
	}
}
func TestMergeSortYH() error {
	data := []int{34,33,12,78,21,1,98,100,50}
	fmt.Println(data)
	mergeSort(data)
	fmt.Println(data)
	return nil
}