package recircution

import "fmt"

// 递推公式，sort(left,mid),sort(mid+1 right
//终止条件 left==right
func mergeSort(data []int)  {
	sort(data,0, len(data) -1 )

}
// 时间复杂度2/n + 2/n +... =   O(n+nlogn) = nlogn
// 空间复杂度
func sort(arr []int, left, right int)  {
	if left == right {
		return
	}
	mid :=  (left + right) / 2
	sort(arr, left, mid)
	sort(arr, mid +1, right)
	merge(arr, left, mid, right)
}
func merge(arr []int, left, mid, right int)  {
	temp := make([]int,right-left + 1)
	i:=left
	j:= mid+1
	tempPos:=0
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
	for tempPos= 0; tempPos< len(temp);tempPos++{
		arr[left] = temp[tempPos]
		left++
	}
}

func TestMergeSort() error {
	data := []int{34,33,12,78,21,1,98,100,50}
	fmt.Println(data)
	mergeSort(data)
	fmt.Println(data)
	return nil
}