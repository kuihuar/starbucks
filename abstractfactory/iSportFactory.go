package abstractfactory

import "fmt"

type iSportFactory interface {
	MakeShoe() iShoe
	MakeShort() iShort
}

func GetSportsFactory(brand string) (iSportFactory, error)  {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
