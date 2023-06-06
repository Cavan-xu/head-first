package decoration

import (
	"fmt"
	"testing"
)

func TestDecoration(t *testing.T) {
	beverage := NewBeverage()
	fmt.Println(beverage.Cost())
	beverage.WithAdditional(WithMilk)
	fmt.Println(beverage.Cost())
	beverage.WithAdditional(WithSoy)
	fmt.Println(beverage.Cost())
}
