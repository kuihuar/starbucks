package validate

import (
	"github.com/asaskevich/govalidator"
	"fmt"
)

type employee1 struct {
	Name string `valid:"required"`
}

func ValidateStruct1 (){
	e := employee1{}
	res,err := govalidator.ValidateStruct(e)
	fmt.Printf("res=%+v, err=%+v", res, err)
}