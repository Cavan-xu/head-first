package pizza

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	nyPizzaStore := NYPizzaStore{
		&PizzaStore{},
	}
	cheesePizza := nyPizzaStore.createPizza("cheese")
	veggiePizza := nyPizzaStore.createPizza("veggie")
	nyPizzaStore.OrderPizza(cheesePizza)
	nyPizzaStore.OrderPizza(veggiePizza)
}
