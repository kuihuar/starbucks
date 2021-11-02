package interviewquestion

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	s := []int{6,7,8}
	fmt.Println("s=",s)
	s1 := SliceAppend(s)
	fmt.Println("s=",s)
	fmt.Println("s1=",s1)
}
