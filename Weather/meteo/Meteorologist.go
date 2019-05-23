package meteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Meteorologist struct{}

func (m *Meteorologist) GetWeather(city string) (w WeatherT) {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyByteSlice, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}

	err = json.Unmarshal(bodyByteSlice, &w)
	if err != nil {
		fmt.Println("error:", err)
	}
	return w
}
