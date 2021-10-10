package search


//单调性（递增或递减）
//存在上下界
//能够通过索引访问

//模板
//left,right = 0,len(arr)-1
//while left <=rigjt:
//	mid = (left+right) / 2
//	if arr[mid] == target:
//		//find the target
//		break or return result
//	elif arr[mid] < target:
//		left + mid+1
//	else:
//		right = mid-1