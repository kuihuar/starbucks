package abstractfactory

import "fmt"

type iShort interface {
	PrintDetail()
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo (logo string) {
	s.logo = logo
}
func(s *short) getLogo() string {
	return s.logo
}
func (s *short)setSize(size int)  {
	s.size = size
}
func (s *short)getSize()int {
	return s.size
}

func (s *short)PrintDetail() {
	fmt.Printf("Logo:%s\n", s.logo)
	fmt.Printf("Size:%d\n", s.size)
}