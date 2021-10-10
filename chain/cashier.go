package chain

import "fmt"

type cashier struct {
	next department
}

func(c *cashier) execute(p *patient) {
	if p.paymentDone {
		//c.next.execute(p)
		return
	}
	fmt.Println("Cashier getting money from patient")
	//p.paymentDone = true
	//c.next.execute(p)
}

func (c *cashier)setNext(next department) {
	c.next = next
}