package command

import "fmt"

type command interface {
	execute()
}

type button struct {
	command command
}

func(b *button) press() {
	b.command.execute()
}
type onCommand struct {
	device device
}
func (on *onCommand) execute() {
	on.device.on()
}
type offCommand struct {
	device device
}

func (off *offCommand)execute()  {
	off.device.off()
}
type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}
func(t *tv) on () {
	t.isRunning = true
	fmt.Println("Turning tv on")
}
func(t *tv) off () {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
func CommandRunTv()  {
	tv := &tv{}
	onCommand := &onCommand{device: tv}
	offCommand:= &offCommand{device: tv}
	onButton := &button{
		command: onCommand,
	}
	offButton := &button{command: offCommand}
	onButton.press()
	offButton.press()
}