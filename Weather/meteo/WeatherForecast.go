package meteo

import (
	"fmt"
	"time"
)

type WeatherForecast struct{}

func (wf *WeatherForecast) FormatWeather(w WeatherT) string {
	windSpd, windGust, windDir := w.GetWind()
	temp, _, _ := w.GetTemperature()

	gustStr := ""
	if windGust > 0 {
		gustStr = fmt.Sprintf(" с порывами до %d м/с", windGust)
	}

	s1 := fmt.Sprintf("Сегодня в городе %s %s, температура воздуха %.0f °C, ветер %s %d м/c%s. Влажность воздуха %d%%. ",
		w.City, w.GetCloudiness(), temp, windDir, windSpd, gustStr, w.GetHumidity())

	s1 += fmt.Sprintf(" Восход солнца %02d:%02d, заход солнца %02d:%02d.",
		time.Unix(w.SysToken.Sunrise, 0).Hour(), time.Unix(w.SysToken.Sunrise, 0).Minute(),
		time.Unix(w.SysToken.Sunset, 0).Hour(), time.Unix(w.SysToken.Sunset, 0).Minute())

	return s1
}
