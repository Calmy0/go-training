package meteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// var response = []byte(`{"coord":{"lon":30.34,"lat":53.91},"weather":[{"id":800,"main":"Clear","description":"ясно","icon":"01d"}],
// "base":"stations","main":{"temp":18.28,"pressure":1012,"humidity":49,"temp_min":16.11,"temp_max":22},
// "visibility":10000,"wind":{"speed":5,"deg":70,"gust":8},
// "clouds":{"all":0},"dt":1558456052,
// "sys":{"type":1,"id":8938,"message":0.0047,"country":"BY","sunrise":1558403263,"sunset":1558461778},
// "id":625665,"name":"Mahilyow","cod":200}`)

type Meteorologist struct{}

func (m *Meteorologist) GetWeather(city string) (w WeatherT) {

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&lang=ru&units=metric&appid=2c19a8c670afc70f2ae7a81f229fce3d")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	BodyByteSlice, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}

	// fmt.Println(string(BodyByteSlice))

	err = json.Unmarshal(BodyByteSlice, &w)
	if err != nil {
		fmt.Println("error:", err)
	}
	return w
}
