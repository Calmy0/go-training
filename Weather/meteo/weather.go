package meteo

type weatherTokenT struct {
	Description string `json:"description"`
}
type mainTokenT struct {
	Temp     float32 `json:"temp"`
	TempMin  float32 `json:"temp_min"`
	TempMax  float32 `json:"temp_max"`
	Humidity int     `json:"humidity"`
}
type windTokenT struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
	Gust  float32 `json:"gust"`
}
type sysTokenT struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}
type WeatherT struct {
	City         string          `json:"name"`
	WeatherToken []weatherTokenT `json:"weather"`
	MainToken    mainTokenT      `json:"main"`
	WindToken    windTokenT      `json:"wind"`
	SysToken     sysTokenT       `json:"sys"`
}

func (w *WeatherT) GetTemperature() (temp, tempMin, tempMax float32) {
	return w.MainToken.Temp, w.MainToken.TempMin, w.MainToken.TempMax
}

func (w *WeatherT) GetCloudiness() (description string) {
	return w.WeatherToken[0].Description
}

func (w *WeatherT) GetHumidity() (humidity int) {
	return w.MainToken.Humidity
}

func (w *WeatherT) GetWind() (speed, gust int, dir string) {
	dirSpec := []struct {
		deg  float32
		name string
	}{
		{0, "северный"},
		{45, "северо-восточный"},
		{90, "восточный"},
		{135, "юго-восточный"},
		{180, "южный"},
		{225, "юго-западный"},
		{270, "западный"},
		{315, "cеверо-западный"},
	}

	for _, ds := range dirSpec {
		if (w.WindToken.Deg >= ds.deg-45/2) &&
			(w.WindToken.Deg <= ds.deg+45/2) {
			dir = ds.name
			break
		}
	}

	return int(w.WindToken.Speed), int(w.WindToken.Gust), dir
}
