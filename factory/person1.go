package factory

type PersonOne struct {
	name string
	age int
}

func NewPersonOneFactory (age int) func(name string) PersonOne {
	return func(name string )PersonOne {
		return PersonOne{
			name: name,
			age: age,
		}
	}
}

