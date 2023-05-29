package structural

import "log"

//

type IPizza interface {
	getDescription() string
	getPrice() int
}

type SmallPizza struct {
}

func (p *SmallPizza) getPrice() int {
	return 300
}

func (p *SmallPizza) getDescription() string {
	return "small pizza"
}

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 30
}

func (t *TomatoTopping) getDescription() string {
	description := t.pizza.getDescription()
	return description + " " + "tomato topping"
}

type CheeseTopping struct {
	pizza IPizza
}

func (t *CheeseTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 50
}

func (t *CheeseTopping) getDescription() string {
	description := t.pizza.getDescription()
	return description + " " + "cheese topping"
}

func RunDecorator() {
	pizza1 := SmallPizza{}
	log.Printf("%s costs %d", pizza1.getDescription(), pizza1.getPrice())

	pizza2 := SmallPizza{}
	cheesePizza := CheeseTopping{pizza: &pizza2}
	log.Printf("%s costs %d", cheesePizza.getDescription(), cheesePizza.getPrice())

	pizza3 := SmallPizza{}
	tomatoPizza := TomatoTopping{pizza: &pizza3}
	log.Printf("%s costs %d", tomatoPizza.getDescription(), tomatoPizza.getPrice())

	pizza4 := SmallPizza{}
	CheeseAndTomatoPizza := CheeseTopping{pizza: &TomatoTopping{pizza: &pizza4}}

	log.Printf("%s costs %d", CheeseAndTomatoPizza.getDescription(), CheeseAndTomatoPizza.getPrice())
}
