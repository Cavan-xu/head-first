package adapter

import (
	"testing"
)

func TestTurkeyAdapter(t *testing.T) {
	var duck IDuck
	adapter := NewTurkeyAdapter(&Turkey{})
	duck = adapter
	duck.Fly()
	duck.Quack()
}
