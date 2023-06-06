package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	c1 := GetInstance()
	c2 := GetInstance()
	t.Log(c1 == c2)
}
