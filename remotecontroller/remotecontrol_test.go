package remotecontrol

import (
	"testing"
)

func TestSimpleRemoteControlPress(t *testing.T) {
	remoteControl := NewSimpleRemoteControl()

	light := &XiaomiLight{}
	lightOnCmd := NewLightOnCommand(light)
	lightOffCmd := NewLightOffCommand(light)
	remoteControl.SetOnCommand(0, lightOnCmd)
	remoteControl.SetOffCommand(0, lightOffCmd)

	door := &GarageDoor{}
	doorUpCmd := NewDoorUpCommand(door)
	doorDownCmd := NewDoorDownCommand(door)
	remoteControl.SetOnCommand(1, doorUpCmd)
	remoteControl.SetOffCommand(1, doorDownCmd)

	remoteControl.OnButtonPress(0)
	remoteControl.OffButtonPress(0)
	remoteControl.OnButtonPress(1)
	remoteControl.OffButtonPress(1)
	remoteControl.UndoButtonPress()
}
