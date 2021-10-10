package builder

type iglooBuilder struct {
	windowType string
	doorType string
	floor int
}


func NewIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}


func (b *iglooBuilder)setWindowType(){
	b.windowType = "Snow Window"
}
func (b *iglooBuilder)setDoorType(){
	b.doorType = "Snow Window"
}
func (b *iglooBuilder)setNumDoor(){
	b.floor = 1
}
func (b *iglooBuilder)GetHouse() house {
	return house{
		windowType: b.doorType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}