package remotecontrol

func NewSimpleRemoteControl() *SimpleRemoteControl {
	soltCount := 6
	remoteController := &SimpleRemoteControl{
		OnCommand:  make([]Command, soltCount),
		OffCommand: make([]Command, soltCount),
	}
	for i := 0; i < soltCount; i++ {
		remoteController.OnCommand[i] = &NoCommand{}
		remoteController.OffCommand[i] = &NoCommand{}
	}
	return remoteController
}

type SimpleRemoteControl struct {
	OnCommand  []Command
	OffCommand []Command
}

func (s *SimpleRemoteControl) SetOnCommand(solt int, command Command) {
	s.OnCommand[solt] = command
}

func (s *SimpleRemoteControl) SetOffCommand(solt int, command Command) {
	s.OffCommand[solt] = command
}

func (s *SimpleRemoteControl) OnButtonPress(solt int) {
	s.OnCommand[solt].Execute()
}

func (s *SimpleRemoteControl) OffButtonPress(solt int) {
	s.OffCommand[solt].Execute()
}
