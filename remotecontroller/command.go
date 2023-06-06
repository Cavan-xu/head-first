package remotecontrol

import (
	"fmt"
)

/*
	命令模式
*/

type Command interface {
	Execute()
}

type NoCommand struct {
}

func (c *NoCommand) Execute() {
}

type LightOnCommand struct {
	light Light
}

func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (l *LightOnCommand) Execute() {
	l.light.On()
}

type LightOffCommand struct {
	light Light
}

func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (l *LightOffCommand) Execute() {
	l.light.Off()
}

type Light interface {
	On()
	Off()
}

type XiaomiLight struct {
}

func (x *XiaomiLight) On() {
	fmt.Println("xiaomi light on")
}

func (x *XiaomiLight) Off() {
	fmt.Println("xiaomi light off")
}

type DoorUpCommand struct {
	door Door
}

func NewDoorUpCommand(door Door) *DoorUpCommand {
	return &DoorUpCommand{door: door}
}

func (d *DoorUpCommand) Execute() {
	d.door.Up()
}

type DoorDownCommand struct {
	door Door
}

func NewDoorDownCommand(door Door) *DoorDownCommand {
	return &DoorDownCommand{door: door}
}

func (d *DoorDownCommand) Execute() {
	d.door.Down()
}

type Door interface {
	Up()
	Down()
}

type GarageDoor struct {
}

func (g *GarageDoor) Up() {
	fmt.Println("garage door on")
}

func (g *GarageDoor) Down() {
	fmt.Println("garage door off")
}
