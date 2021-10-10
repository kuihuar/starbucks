package abstractfactory


type adidas struct {

}

func(a *adidas) MakeShoe () iShoe{
	return &adidasShoe{
		shoe: shoe{
			logo:"adidas",
			size: 14,
		},
	}
}
func (a *adidas) MakeShort() iShort {
	return &adidasShort{struct {
		logo string
		size int
	}{logo: "adidas", size: 14}}
}