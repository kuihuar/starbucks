package array_style

//func MaxArea(height []int) int {
//	var max, area int
//	//双重遍历i,j不会重复
//	for i:=0; i< len(height)-1; i++{
//		for j:=i+1; j< len(height); j++{
//			//fmt.Printf("i=%d * j=%d\n", i,j)
//			if height[i] > height[j] {
//				area = (j-i) * height[j]
//			} else {
//				area = (j-i) * height[i]
//			}
//			if area > max {
//				max = area
//			}
//		}
//	}
//	return max
//}
func count()  {

}
//左右边界向中靠拢
func MaxArea(height []int) int {
	var max,area int
	for l,r:=0,len(height)-1; l<r;{
		if height[l] < height[r]{
			area = (r - l) * height[l]
			l++
		}else{
			area = (r - l) * height[r]
			r--
		}
		if area > max {
			max = area
		}
	}
	return max
}