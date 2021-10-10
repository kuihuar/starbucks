package bridge

import "fmt"

type computer interface {
	print()
	setPrinter(printer)
}

type printer interface {
	printFile()
}
type mac struct {
	printer printer
}
func(m *mac) print(){
	fmt.Println("print request for mac")
	m.printer.printFile()
}
func(m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func(w *windows) print(){
	fmt.Println("print request for windows")
	w.printer.printFile()
}
func(w *windows) setPrinter(p printer) {
	w.printer = p
}

type epson struct {
}
func (p *epson) printFile(){
	fmt.Println("Printing by a EPSON printer")
}
type hp struct {
}
func (h *hp) printFile() {
	fmt.Println("Printing by a HP printer")
}

func TestBridge(){
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println("+++++++++++++")

	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}