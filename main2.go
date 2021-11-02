package main

import "fmt"

func init()  {
	fmt.Println("bar")
}
func init()  {
	fmt.Println("bar1")
}

// 最多六位正整数，保证每一位不重复，且必须包含42的数有多少
var visited map[int]bool
var res [][]int

type OtherInterface interface {
	OtherFunc()
}

type Printer interface {
	SomePrint()
}
type user struct {
	name string
	age int
}

func (u user)SomePrint()  {
	fmt.Printf("%+v\n",u)
}

func (u user)OtherFunc()  {
	fmt.Printf("other..%+v\n",u)
}
func consumer(data chan int, done chan bool)  {
	for x := range data {
		println("recv:",x)
	}
	done <- true
}
func producer(data chan int)  {
	for i:=0; i<4;i++{
		println("send:",i)
		data <-i
	}
	close(data)
}
var g = 100

func test() *int {
	a := 0x100
	return &a
}
func main()  {
var a *int = test()
println(a, *a)



	//var ages map[string]int
	//ages = make(map[string]int)
	//
	//ages2 := make(map[string]int)
	//
	//ages3 := map[string]int{
	//	"alice":33,
	//}
	//channelmap := make(map[chan int]int)
	//ch := make(chan int)
	//channelmap[ch] = 7
	//fmt.Println(channelmap)
	//ch2 := make(chan int,2)
	//channelmap[ch2] = 3
	//fmt.Println(channelmap )
	////delete(ages, "alice")
	////fmt.Println(ages["alice"])
	//fmt.Println(ages,ages2, ages3)
	//visited = map[int]bool{}
	//nums := [10]int{0,1,2,3,4,5,6,7,8,9}
	//maxLen:= 6
	//for i:=0; i<10;i++{
	//	dfs(nums, i, []int{},10, maxLen)
	//}
	//fmt.Printf("共有%v位",len(res))
}
func dfs(nums [10]int, index int, curItem []int, length int, maxlen int){
	if visited[index] {
		return
	}
	visited[index] = true
	curItem = append(curItem, nums[index])
	if len(curItem) == maxlen  && (
		(curItem[0] == 4 && curItem[1]==2) ||
			(curItem[0] != 0 && curItem[1] == 4 && curItem[2]==2) ||
			(curItem[0] != 0 && curItem[2] == 4 && curItem[3]==2)||
			(curItem[0] != 0 && curItem[3] == 4 && curItem[4]==2)||
			(curItem[0] != 0 && curItem[4] == 4 && curItem[5]==2)){
		cur := make([]int,maxlen)
		copy(cur, curItem[0:maxlen])
		//fmt.Printf("%v\n", cur)
		res = append(res, cur)
	}
	for i:=0; i<length;i++{
		dfs(nums, i, curItem,10,maxlen)
	}
	curItem = curItem[0:len(curItem)-1]
	visited[index] = false
}