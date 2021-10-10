package errors

import (
	"errors"
	"fmt"
)

func handleErrors (){
	inlineErr := errors.New("inline error occured")

	outErr := fmt.Errorf("outline error occured %w",inlineErr)
	returnErr := fmt.Errorf("request error occured %w", outErr)
	fmt.Println(returnErr)
	fmt.Println(errors.Unwrap(returnErr))
	fmt.Println(errors.Unwrap(outErr))
	fmt.Println(errors.Unwrap(inlineErr))
	fmt.Println(inlineErr == outErr)
	fmt.Println(errors.Is(inlineErr, outErr))
	fmt.Println(errors.As(inlineErr, outErr))
}
