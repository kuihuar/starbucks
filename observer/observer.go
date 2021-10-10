package observer

import "fmt"

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

type item struct {
	observerList []observer
	name string
	inStock bool
}

func newItem(name string) *item {
	return &item{name: name}
}
func (i *item)register(ob observer)  {
	i.observerList = append(i.observerList,ob)
}
func (i *item)deregister(ob observer)  {
	i.observerList = removerFromSlice(i.observerList,ob)
}
func removerFromSlice(observerList []observer,observerRemove observer)[]observer{
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observer.getID() == observerRemove.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
func (i *item)notifyAll()  {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
func (i *item)updateAvailability()  {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
type observer interface {
	update(string2 string)
	getID()string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n",c.id, itemName)
}
func (c *customer)getID() string {
	return c.id
}

func TestObserver()  {
	shirtItem := newItem("Nike shirt")
	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
}