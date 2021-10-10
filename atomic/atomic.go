package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var config atomic.Value
var lock sync.Mutex
func TestConfig()  {
	var scene * sync.Map = new (sync.Map)
	scene.Store("aaa", 97)
	vlue,ok := scene.Load("aaa")
	fmt.Println(vlue,ok)
	scene.Delete("aaa")
	scene.LoadOrStore("aaa", 98)
	scene.Range(func(key, value interface{}) bool {
		if key == "aaa" {

		}
		return true
	})
}
func getConfig() map[string]string {
	if config.Load() == nil {
		fmt.Println("init config")
		config.Store(map[string]string{})
		return config.Load().(map[string]string)
	}
	return config.Load().(map[string]string)
}
