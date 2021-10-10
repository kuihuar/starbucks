package memento

import "fmt"

type originator struct {
	state string
}

func (e *originator) creatememento() *memento {
	return &memento{
		state: e.state,
	}
}

func (e *originator)restorememento(m *memento)  {
	e.state = m.getSavedState()
}

func (e *originator) setState(state string)  {
	e.state = state
}
func (e *originator) getState() string {
	return e.state
}
type memento struct {
	state string
}

func (m *memento)getSavedState() string {
	return m.state
}

type caretaker struct {
	mementoSlice []*memento
}

func (c *caretaker)addMemento(m *memento)  {
	c.mementoSlice = append(c.mementoSlice,m)
}
func (c *caretaker)getmemento(index int) *memento {
	return c.mementoSlice[index]
}

func TestMemento()  {
	caretaker := &caretaker{mementoSlice: make([]*memento, 0)}

	originator := &originator{state: "A"}
	fmt.Printf("originator current state:%s\n",originator.getState())
	caretaker.addMemento(originator.creatememento())

	originator.setState("B")
	fmt.Printf("originator current state:%s\n",originator.getState())
	caretaker.addMemento(originator.creatememento())
	originator.setState("C")
	fmt.Printf("originator current state:%s\n",originator.getState())
	caretaker.addMemento(originator.creatememento())

	originator.restorememento(caretaker.getmemento(1))
	fmt.Printf("originator current state:%s\n",originator.getState())
	originator.restorememento(caretaker.getmemento(0))
	fmt.Printf("originator current state:%s\n",originator.getState())
}