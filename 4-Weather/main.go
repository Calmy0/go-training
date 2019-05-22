package main

import (
	"fmt"
	"weater/meteo"
)

func main() {

	var input string

	print("Укажите город:")

	_, err := fmt.Scanln(&input)
	if err != nil {
		print("Error while reading.\n")
		return
	}

	m := new(meteo.Meteorologist)
	wf := new(meteo.WeatherForecast)

	fmt.Println(wf.FormatWeather(m.GetWeather(input)))

}
