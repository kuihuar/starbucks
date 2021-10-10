package abstractfactory

type nike struct {

}

func(a *nike) MakeShoe () iShoe{
	return &nikeShoe{
		shoe: shoe{
			logo:"nike",
			size: 14,
		},
	}
}
func (a *nike) MakeShort() iShort {
	//return &nikeShort{struct {
	//	logo string
	//	size int
	//}{logo: "nike", size: 14}}
	return &nikeShort{short{
		logo: "nike",
		size: 14,
	}}
}
