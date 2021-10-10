package jkgo

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + ": not found"
}

func Test()  {
	_, err := os.Open("abcd")
	if e, ok := err.(*NotFoundError); ok {
		fmt.Println(e)
	}

}

type QueryError struct {
	Query string
	Err error
}

func IsAs (){
	var err  error = nil
	var ErrNotFound error = errors.New("not found")
	if err == ErrNotFound {

	}
	if errors.Is(err, ErrNotFound){
		//soemthing not found
	}
	var e *QueryError
	if errors.As(err,&e) {
		//err is a *queryError, and e is set the erros's value
	}
	fmt.Errorf("decompress %v: %w", " decompress name", err)

	err = fmt.Errorf("acess denied: %w", ErrNotFound)
	if errors.Is(err, ErrNotFound) {

	}
}
