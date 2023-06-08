package template

import (
	"fmt"
)

/*
	模板模式
*/

type ICaffeineBeverage interface {
	prepareRecipe(ICaffeineBeverage)
	boilWater()
	brew()
	addCondition()
	pourInCup()
}

type CaffeineBeverage struct {
}

func (c *CaffeineBeverage) prepareRecipe(beverage ICaffeineBeverage) {
	c.boilWater()
	beverage.brew()
	beverage.addCondition()
	c.pourInCup()
}

func (c *CaffeineBeverage) boilWater() {
}

func (c *CaffeineBeverage) brew() {
	panic("implement me")
}

func (c *CaffeineBeverage) addCondition() {
	panic("implement me")
}

func (c *CaffeineBeverage) pourInCup() {
}

type CoffeeBeverage struct {
	*CaffeineBeverage
}

func NewCoffeeBeverage(c *CaffeineBeverage) *CoffeeBeverage {
	return &CoffeeBeverage{c}
}

func (c *CoffeeBeverage) brew() {
	fmt.Println("brew coffee")
}

func (c *CoffeeBeverage) addCondition() {
	fmt.Println("add sugar in coffee")
}

type TeaBeverage struct {
	*CaffeineBeverage
}

func NewTeaBeverage(c *CaffeineBeverage) *TeaBeverage {
	return &TeaBeverage{c}
}

func (c *TeaBeverage) brew() {
	fmt.Println("brew tea")
}

func (c *TeaBeverage) addCondition() {
	fmt.Println("add tea leaves")
}
