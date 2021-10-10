package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
type employee struct {
	Name string `validate:"required"`
}

func ValidateStruct (){
	e := employee{}

	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		fmt.Printf("Error :%s\n", err)
	}
}