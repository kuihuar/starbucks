package search

import (
	"reflect"
)

func minMutation(start string, end string, bank []string) int {
	res := 0
	for _, bankItem := range bank {
		if reflect.DeepEqual(end, bankItem) {
			res ++
		}
	}
	if res > 0 {
		return -1
	}
return 1

}
