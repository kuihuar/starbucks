package factory

import "fmt"

func SomeAnimal()  {
	docFactory := AnimalFactory{
		species:"dog",
		houseName:"kernel",
	}
	dog := docFactory.NewAnimal(2)
	kennel := docFactory.NewHouse(3)




	horseFactory := AnimalFactory{
		species: "horse",
		houseName: "stable",
	}
	fmt.Println(dog,kennel)
	horse := horseFactory.NewAnimal(12)
	stable := horseFactory.NewHouse(30)
	fmt.Println(horse,stable)


}
func SomePersonOne(){
	newBaby := NewPersonOneFactory(1)
	baby:= newBaby("john")

	fmt.Println(baby.name)
	newTeenager := NewPersonOneFactory(16)
	teen := newTeenager("jill")
	fmt.Println(teen.name)
}
