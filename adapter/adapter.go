package adapter

import (
	"fmt"
)

/*
	适配器模式
*/

type IDuck interface {
	Quack()
	Fly()
}

type Turkey struct {
}

func (t *Turkey) Gobble() {
	fmt.Println("turkey gobble")
}

func (t *Turkey) Fly() {
	fmt.Println("turkey fly")
}

type TurkeyAdapter struct {
	turkey *Turkey
}

func NewTurkeyAdapter(t *Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{
		turkey: t,
	}
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	t.turkey.Fly()
}
