package pizza

import (
	"fmt"
)

type PizzaStore struct {
}

func (c *PizzaStore) OrderPizza(pizza Pizza) {
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
}

func (c *PizzaStore) createPizza(name string) Pizza {
	return nil
}

type NYPizzaStore struct {
	*PizzaStore
}

func (c *NYPizzaStore) createPizza(name string) Pizza {
	var pizza Pizza
	switch name {
	case "cheese":
		pizza = &NYCheesePizza{}
	case "veggie":
		pizza = &NYVeggiePizza{}
	}
	return pizza
}

type NYCheesePizza struct {
}

func (c *NYCheesePizza) prepare() {
	fmt.Println("prepare NY cheese pizza")
}

func (c *NYCheesePizza) bake() {
}

func (c *NYCheesePizza) cut() {
}

func (c *NYCheesePizza) box() {
}

type NYVeggiePizza struct {
}

func (v *NYVeggiePizza) prepare() {
	fmt.Println("prepare NY veggie pizza")
}

func (v *NYVeggiePizza) bake() {
}

func (v *NYVeggiePizza) cut() {
}

func (v *NYVeggiePizza) box() {
}
