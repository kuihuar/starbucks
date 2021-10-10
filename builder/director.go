package builder

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{builder: b}
}

func (d *director)BuildHouse() house{
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumDoor()
	return d.builder.GetHouse()
}