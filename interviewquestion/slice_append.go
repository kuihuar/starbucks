package interviewquestion

func SliceAppend(s []int) []int {
	for i:=0; i<10;i++{
		s=append(s,i)
	}
	//s[2]=100
	return s
}