package cond

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func listen(name string , a map[string]int, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println(name,  " age:", a["T"])
	c.L.Unlock()
}

func broadCast(name string, a map[string]int, c *sync.Cond)  {
	time.Sleep(time.Second)
	c.L.Lock()
	a["T"] = 25
	c.Broadcast()
	c.L.Unlock()
}

func TestCond()  {
	var age = make(map[string]int)
	m:= sync.Mutex{}
	cond := sync.NewCond(&m)
	go listen("lis1", age, cond)
	go listen("lis2", age, cond)
	go broadCast("lis1", age, cond)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

}