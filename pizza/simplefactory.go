package pizza

import (
	"fmt"
)

type Pizza interface {
	prepare()
	bake()
	cut()
	box()
}

type SimplePizzaStore struct {
	factory *PizzaFactory
}

func (c *SimplePizzaStore) createPizza(name string) {
	pizza := c.factory.NewPizza(name)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
}

type PizzaFactory struct {
}

func (c *PizzaFactory) NewPizza(name string) Pizza {
	var pizza Pizza
	switch name {
	case "cheese":
		pizza = &CheesePizza{}
	case "veggie":
		pizza = &VeggiePizza{}
	}
	return pizza
}

type CheesePizza struct {
}

func (c *CheesePizza) prepare() {
	fmt.Println("prepare cheese pizza")
}

func (c *CheesePizza) bake() {
}

func (c *CheesePizza) cut() {
}

func (c *CheesePizza) box() {
}

type VeggiePizza struct {
}

func (v *VeggiePizza) prepare() {
	fmt.Println("prepare veggie pizza")
}

func (v *VeggiePizza) bake() {
}

func (v *VeggiePizza) cut() {
}

func (v *VeggiePizza) box() {
}
