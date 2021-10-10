package adapter

import "fmt"

type computer interface {
	insertInSquarePort()
}

type mac struct {

}
func (m *mac) insertInSquarePort () {
	fmt.Println("insert square port into mac machine")
}

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter)insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

type windows struct {}

func (w *windows)insertInCirclePort()  {
	fmt.Println("insert circle port into windows machine")
}
type client struct {

}
func (c *client)insertSquareUsbInComputer(com computer){
	com.insertInSquarePort()
}

func TestAdapter()  {
	client := &client{}

	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}