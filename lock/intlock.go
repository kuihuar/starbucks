package lock

import "sync"

type intLock struct {
	val int
	sync.Mutex
}

func NewLock(val int) *intLock {
	return &intLock{
		val: val,
	}
}
func ( lock *intLock)GetVal() int{
	return lock.val
}
func ( lock *intLock)IncVal() {
	lock.val++
}
func (n *intLock) IsEven() bool{
	return n.val %2 == 0
}