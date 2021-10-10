package lock

import (
	"fmt"
	"testing"
	"time"
)
func TestIsEven(t *testing.T)  {
	n:= &intLock{val: 0}



	go func() {
		n.Lock()
		defer n.Unlock()
		nIsEven := n.IsEven()
		time.Sleep(5 * time.Millisecond)
		if nIsEven{
			fmt.Println(n.val, "is even")
			return
		}
		fmt.Println(n.val, "is odd")
	}()

	go func() {
		n.Lock()
		n.val++
		n.Unlock()
	}()

	time.Sleep(time.Second)
}