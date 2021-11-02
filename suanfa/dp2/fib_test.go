package dp2

import "testing"

func TestFib(t *testing.T) {
	n:= Fib1(4)
	t.Log(n)
	t.Fail()
}