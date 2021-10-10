package visitor

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}
type square struct {
	side int
}
func(s *square) accept (v visitor) {
	v.visitForSquare(s)
}
func (s *square)getType() string {
	return "Square"
}
type circle struct {
	radius int
}
func (c *circle)accept (v visitor) {
	v.visitForCircle(c)
}
func (c *circle)getType() string {
	return "Circle"
}

type visitor interface {
	visitForSquare(square2 *square)
	visitForCircle(circle2 *circle)
	visitForRectangle(rectangle2 *rectangle)
}
type rectangle struct {
	l int
	b int
}

func (r *rectangle)accept(v visitor)  {
	v.visitForRectangle(r)
}
func (r *rectangle)getType() string {
	return "Rectangle"
}
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square)  {
	fmt.Println("Calculating area for square")
}
func (a *areaCalculator)visitForCircle(c *circle)  {
	fmt.Println("Calculating area for circle")
}
func(a *areaCalculator)visitForRectangle(r *rectangle){
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x,y int
}

func (m *middleCoordinates) visitForSquare(s *square)  {
	fmt.Println("Calculating area for square")
}
func (m *middleCoordinates)visitForCircle(c *circle)  {
	fmt.Println("Calculating area for circle")
}
func(m *middleCoordinates)visitForRectangle(r *rectangle){
	fmt.Println("Calculating area for rectangle")
}

func TestVisitor (){
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2,b: 3}
	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println("=====")
	middleCoordinates := &middleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
