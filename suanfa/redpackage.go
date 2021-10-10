package suanfa

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePackage(totalRmb, size, minRmb int,  packages *[]int)  {
	if size <=0{
		return
	}
	var newRmb int
	if size == 1 {
		newRmb =  totalRmb
	}else {
		rand.Seed(time.Now().UnixNano())
		newRmb = rand.Intn(totalRmb / size * 2 - minRmb) + minRmb
	}
	*packages = append(*packages, newRmb)
	generatePackage(totalRmb-newRmb, size-1, minRmb,packages)
}

func TestRedPackage()  {

	for i := 0;i< 40;i++ {
		var packages []int
		generatePackage(500, 10,1,&packages)
		fmt.Println(packages)
	}

}