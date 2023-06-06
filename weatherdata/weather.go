package weatherdata

/*
	观察者模式
*/

import (
	"fmt"
)

type WeatherData struct {
}

func (w *WeatherData) getTemperature() int {
	return 0
}

func (w *WeatherData) getHumidity() int {
	return 0
}

func (w *WeatherData) getPressure() int {
	return 0
}

// 任何时候气象测试被更新，将调用此方法
func (w *WeatherData) measureChanges() {
	GetMessageManager().Receive(&Message{
		Temperature: w.getTemperature(),
		Humidity:    w.getHumidity(),
		Pressure:    w.getPressure(),
	})
}

type Message struct {
	Temperature int
	Humidity    int
	Pressure    int
}

var messageManager *MessageManager

func NewMessageManager() *MessageManager {
	messageManager = &MessageManager{
		MessageChan: make(chan *Message, 1024),
	}
	go messageManager.Work()
	return messageManager
}

func GetMessageManager() *MessageManager {
	return messageManager
}

type MessageManager struct {
	MessageChan chan *Message
	ReceiverArr []MessageReceiver
}

func (m *MessageManager) Receive(message *Message) {
	m.MessageChan <- message
}

func (m *MessageManager) Work() {
	for {
		message, ok := <-m.MessageChan
		if !ok {
			return
		}
		for _, receiver := range m.ReceiverArr {
			receiver.Receive(message)
		}
	}
}

func (m *MessageManager) AddReceiver(receiver MessageReceiver) {
	m.ReceiverArr = append(m.ReceiverArr, receiver)
}

func (m *MessageManager) RemoveReceiver(receiver MessageReceiver) {
	for i := 0; i < len(m.ReceiverArr); i++ {
		if receiver.UniqueId() == m.ReceiverArr[i].UniqueId() {
			m.ReceiverArr = append(m.ReceiverArr[:i], m.ReceiverArr[i+1:]...)
			break
		}
	}
}

type MessageReceiver interface {
	Receive(message *Message)
	UniqueId() int
}

type CurrentConditionReceiver struct {
}

func (c *CurrentConditionReceiver) Receive(message *Message) {
	fmt.Println("CurrentConditionReceiver do something")
}

func (c *CurrentConditionReceiver) UniqueId() int {
	return 0
}

type StatisticsDisplayReceiver struct {
}

func (s *StatisticsDisplayReceiver) Receive(message *Message) {
	fmt.Println("StatisticsDisplayReceiver do something")
}

func (c *StatisticsDisplayReceiver) UniqueId() int {
	return 0
}

type ForecastDisplayReceiver struct {
}

func (f *ForecastDisplayReceiver) Receive(message *Message) {
	fmt.Println("ForecastDisplayReceiver do something")
}

func (c *ForecastDisplayReceiver) UniqueId() int {
	return 0
}
