package duck

/*
	策略模式
*/

import (
	"fmt"
)

type FlyBehavior interface {
	Fly()
}

type PerformFlyV1 struct {
}

func (p *PerformFlyV1) Fly() {
	fmt.Println("fly with v1")
}

type PerformFlyV2 struct {
}

func (p *PerformFlyV2) Fly() {
	fmt.Println("fly with v2")
}

type QuackBehavior interface {
	Quack()
}

type PerformQuack struct {
}

func (p PerformQuack) Quack() {
	fmt.Println("quack")
}

type Duck struct {
	FlyBehavior
	QuackBehavior
}

func NewDuck(flyBehavior FlyBehavior, quackBehavior QuackBehavior) *Duck {
	return &Duck{
		FlyBehavior:   flyBehavior,
		QuackBehavior: quackBehavior,
	}
}

func (d *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	d.FlyBehavior = flyBehavior
}

type MallardDuck struct {
	*Duck
}

func NewMallardDuck(flyBehavior FlyBehavior, quackBehavior QuackBehavior) *MallardDuck {
	return &MallardDuck{
		Duck: NewDuck(flyBehavior, quackBehavior),
	}
}
