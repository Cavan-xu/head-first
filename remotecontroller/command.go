package remotecontrol

import (
	"fmt"
)

/*
	命令模式
*/

type Command interface {
	Execute()
	Undo()
}

type NoCommand struct {
}

func (c *NoCommand) Execute() {
}

func (c *NoCommand) Undo() {
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

func (l *LightOnCommand) Undo() {
	l.light.Off()
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

func (l *LightOffCommand) Undo() {
	l.light.On()
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

func (d *DoorUpCommand) Undo() {
	d.door.Down()
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

func (d *DoorDownCommand) Undo() {
	d.door.Up()
}

type Door interface {
	Up()
	Down()
}

type GarageDoor struct {
}

func (g *GarageDoor) Up() {
	fmt.Println("garage door up")
}

func (g *GarageDoor) Down() {
	fmt.Println("garage door down")
}
