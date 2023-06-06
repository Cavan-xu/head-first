package duck

import (
	"testing"
)

func TestNewMallardDuck(t *testing.T) {
	flyV1 := &PerformFlyV1{}
	flyV2 := &PerformFlyV2{}
	mallardDuck := NewMallardDuck(flyV1, &PerformQuack{})
	mallardDuck.Fly()
	mallardDuck.SetFlyBehavior(flyV2)
	mallardDuck.Fly()
	mallardDuck.Quack()
}
