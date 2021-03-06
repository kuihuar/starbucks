package factory

type Animal struct {
	species string
	age     int
}

type AnimalHouse struct {
	name         string
	sizeInMeters int
}

type AnimalFactory struct {
	species   string
	houseName string
}

// NewAnimal is an `Animal` factory that fixes
// the species as the value of its `AnimalFactory` instance
func (af AnimalFactory) NewAnimal(age int) Animal{
	return Animal{
		species: af.species,
		age:     age,
	}
}


// NewHouse is an `AnimalHouse` factory that fixes
// the name as the value of its `AnimalFactory` instance
func (af AnimalFactory) NewHouse(sizeInMeters int) AnimalHouse{
	return AnimalHouse{
		name:         af.houseName,
		sizeInMeters: sizeInMeters,
	}
}

