package builder

type normalBuilder struct {
	windowType string
	doorType string
	floor int
}

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder)setWindowType(){
	b.windowType = "Wooden Window"
}
func (b *normalBuilder)setDoorType(){
	b.doorType = "Wooden Door"
}
func (b *normalBuilder)setNumDoor(){
	b.floor = 2
}
func (b *normalBuilder)GetHouse() house {
	return house{
		windowType: b.doorType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}