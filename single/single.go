package single

import (
	"fmt"
	"sync"
)

type single struct {

}
var lock = &sync.Mutex{}
var singleInstance *single
var once sync.Once
func GetInstance () *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("creating Single instance Now.")
			singleInstance = &single{}
		} else {
			fmt.Println("created-1")
		}
	} else {
		fmt.Println("created-2")
	}
	return singleInstance
}

func GetInstance2() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("creating single instance now..")
			singleInstance = &single{}
		})
	}else {
		fmt.Println("created-2")
	}
	return singleInstance
}