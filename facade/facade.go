package facade

import (
	"fmt"
)

/*
	外观模式
*/

func NewHomeTheaterFacade(tuner *Tuner, screen *Screen, light *Light) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tuner:  tuner,
		screen: screen,
		light:  light,
	}
}

type HomeTheaterFacade struct {
	tuner  *Tuner
	screen *Screen
	light  *Light
}

func (h *HomeTheaterFacade) watchMovie() {
	h.tuner.On()
	h.screen.Up()
	h.screen.On()
	h.light.On()
}

func (h *HomeTheaterFacade) endMovie() {
	h.tuner.Off()
	h.screen.Off()
	h.light.Off()
}

type Tuner struct {
}

func (t *Tuner) On() {
	fmt.Println("tuner on")
}

func (t *Tuner) Off() {
	fmt.Println("tuner off")
}

func (t *Tuner) SetAm() {
}

type Screen struct {
}

func (s *Screen) On() {
	fmt.Println("screen on")
}

func (s *Screen) Off() {
	fmt.Println("screen off")
}

func (s *Screen) Up() {
	fmt.Println("screen up")
}

func (s *Screen) Down() {
	fmt.Println("screen down")
}

type Light struct {
}

func (l *Light) On() {
	fmt.Println("light on")
}

func (l *Light) Off() {
	fmt.Println("light off")
}
