package factorygun

type maverick struct {
	gun
}

func newMoverick ()iGun {
	return &maverick{gun{
		name:  "maverick gun",
		power: 5,
	}}
}