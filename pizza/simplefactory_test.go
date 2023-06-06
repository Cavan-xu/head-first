package pizza

import (
	"testing"
)

func TestNewSimpleFactory(t *testing.T) {
	pizzaStore := &SimplePizzaStore{
		factory: &PizzaFactory{},
	}
	pizzaStore.createPizza("cheese")
	pizzaStore.createPizza("veggie")
}
