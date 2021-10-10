package abstractfactory

import "fmt"

type iShoe interface {
	PrintDetail()
}

type shoe struct {
	logo  string
	size int
}

func (s *shoe) setLogo (logo string) {
	s.logo = logo
}
func(s *shoe) getLogo() string {
	return s.logo
}
func (s *shoe)setSize(size int)  {
	s.size = size
}
func (s *shoe)getSize()int {
	return s.size
}
func (s *shoe)PrintDetail() {
	fmt.Printf("Logo:%s\n", s.logo)
	fmt.Printf("Size:%d\n", s.size)
}