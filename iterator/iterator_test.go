package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	var iterator Iterator
	iterator = NewPancakeHouseMenu()
	for iterator.HasNext() {
		t.Log(iterator.Next().GetName())
	}

	iterator = NewDinnerMenu()
	for iterator.HasNext() {
		t.Log(iterator.Next().GetName())
	}
}
