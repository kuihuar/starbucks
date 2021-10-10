package recircution

//找一个分区点， 大于分区点的放到右边， 小于分区点的放左边
//子问题和大问题的求解方式一样
//最后左右子数只剩下一个元素，就没有必要分了
//存在最小子部碮同，子数组只剩下一个元素或为空

//递推公式和终止条件

//极端情况下，退化至 O(n^2)
//如何避免
// 1 三数取中
// 2 随机法

func QuickSort(data []int) []int {
	quickSort(data,0, len(data)-1)
	return data
}

func quickSort(data []int, low, hight int)  {

	if low >= hight {
		return
	}
	pivot := partition(data, low, hight)
	quickSort(data, low, pivot-1)
	quickSort(data, pivot + 1, hight)
}
//0		1		2		3		4		5		6		7		8		9		10	11		low pivot	hight
//12	9		34	22	23	13	14	77	45	8		10	20
//12	9		13	14 	8		10	20	34	22	33	77	45		0		6		11
//9		8		10	12	13	14	20	34	22	33	77	45		0		2 	5
//8		9		10	12	13	14	20	34	22	33	77	45		0		0		1
//8		9		10	12	13	14	20	34	22	33	77	45		1				1
//8		9		10	12	13	14	20	34	22	33	77	45		3		5		5
//8		9		10	12	13	14	20	34	22	33	77	45		3		4		4
//8		9		10	12	13	14	20	34	22	33	45	77		7		10	11
//8		9		10	12	13	14	20	22	33	34	45	77		7				7
//8		9		10	12	13	14	20	22	33	34	45	77		9				9
//8		9		10	12	13	14	20	22	33	34	45	77		11			11




func partition(data []int, low, hight int)  int {
	pivot := hight
	less := low
	//great := low
	for great := low;great <= hight-1;great++{ // hight-1 ,hight 位置就是 pivot
		if data[great] < data[pivot] {
			data[less], data[great] = data[great], data[less]
			less++
		}
	}
	data[less],data[hight]= data[hight],data[less]
	return less
}

// 处理大量重复元素的办法
//三路快排，三向切分法， 小于的放一部分，等于的放一起，大于的放一起
//保证以下四点
//[low...less] < pivot
//[less...i) = pivot //不包括i
//[i...great] 未处理
//(great ...hight] > pivot
// 所谓三路，其实是三段，less great i

