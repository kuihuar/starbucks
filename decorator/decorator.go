package decorator

import "fmt"

type pizza interface {
	getPrice() int
}

type peppyPaneer struct {
}

func (p *peppyPaneer) getPrice() int {
	return 20
}

type veggieMania struct {
}

func (p *veggieMania) getPrice() int {
	return 15
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

type tomatoTopping struct {
	pizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func TestDecorator() {
	veggiePizza := &veggieMania{}

	viggiePizzaWithCheese := &cheeseTopping{pizza: veggiePizza}
	viggiePizzaWithCheeseAndTomato := &tomatoTopping{pizza: viggiePizzaWithCheese}

	fmt.Printf("Price of veggieMania prizza with tomato and cheese topping is %d\n", viggiePizzaWithCheeseAndTomato.getPrice())

	peppyPaneerPizza := &peppyPaneer{}
	peppyPaneerPizzaWithCheese := &cheeseTopping{pizza: peppyPaneerPizza}
	fmt.Printf("Price of peppyPaneer prizza with cheese topping is %d\n", peppyPaneerPizzaWithCheese.getPrice())


}
