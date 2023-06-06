package pizza

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	factory := &NYIngredientFactory{}
	cheese := factory.CreateCheese()
	t.Log(cheese.GetCheeseId())
}
