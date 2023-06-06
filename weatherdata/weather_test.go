package weatherdata

import (
	"testing"
)

func TestWeatherDataWorker(t *testing.T) {
	NewMessageManager()
	currentConditionReceiver := &CurrentConditionReceiver{}
	statisticsDisplayReceiver := &StatisticsDisplayReceiver{}
	forecastDisplayReceiver := &ForecastDisplayReceiver{}
	GetMessageManager().AddReceiver(currentConditionReceiver)
	GetMessageManager().AddReceiver(statisticsDisplayReceiver)
	GetMessageManager().AddReceiver(forecastDisplayReceiver)
	weatherData := &WeatherData{}
	for i := 0; i < 10; i++ {
		weatherData.measureChanges()
	}
}
