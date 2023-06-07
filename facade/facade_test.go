package facade

import (
	"testing"
)

func TestHomeTheaterFacade(t *testing.T) {
	control := NewHomeTheaterFacade(&Tuner{}, &Screen{}, &Light{})
	control.watchMovie()
	control.endMovie()
}
