package template

import (
	"testing"
)

func TestTemplates(t *testing.T) {
	coffeeBeverage := NewCoffeeBeverage(&CaffeineBeverage{})
	teaBeverage := NewTeaBeverage(&CaffeineBeverage{})
	coffeeBeverage.prepareRecipe(coffeeBeverage)
	teaBeverage.prepareRecipe(teaBeverage)
}
