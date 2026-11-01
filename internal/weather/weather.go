package weather

import (
	"log"

	"github.com/KillAllChickens/gust/internal/convert"
	"resty.dev/v3"
)

type respDataCurrent struct {
	Time        string  `json:"time"`
	Temp        float64 `json:"temperature_2m"`
	FeelsLike   float64 `json:"apparent_temperature"`
	HumidityPct int     `json:"relative_humidity_2m"`
	WindKph     float64 `json:"wind_speed_10m"`
	WeatherCode int     `json:"weather_code"`
}

type respData struct {
	Current respDataCurrent `json:"current"`
}

func FetchWeather() Forecast {
	lat, long := "37.9967", "-88.9201"
	client := resty.New()
	client.SetBaseURL("https://api.open-meteo.com")
	client.SetHeader("Accept-Encoding", "gzip")
	client.SetQueryParams(map[string]string{
		"current":   "temperature_2m,apparent_temperature,relative_humidity_2m,wind_speed_10m,weather_code",
		"hourly":    "temperature_2m,wind_speed_10m,weather_code,precipitation_probability",
		"daily":     "temperature_2m_max,temperature_2m_min,precipitation_probability_max,weather_code,sunrise,sunset",
		"timezone":  "auto",
		"latitude":  lat,
		"longitude": long,
	})

	var finalData respData
	forecastResp, err := client.R().
		SetResult(&finalData).
		Get("/v1/forecast")
	if err != nil {
    log.Fatalf("request failed: %v", err)
	}
	if forecastResp.StatusCode() >= 400 {
    log.Fatalf("bad status %d: %s", forecastResp.StatusCode(), forecastResp.String())
	}
	var condition Condition

	switch finalData.Current.WeatherCode {
	case 0, 1:
		condition = Clear
	case 2, 3:
		condition = Cloudy
	case 45, 48:
		condition = Fog
	case 51, 53, 55, 56, 57, 61, 63, 65, 66, 67, 80, 81, 82:
		condition = Rain
	case 71, 73, 75, 77, 85, 86:
		condition = Snow
	case 95, 96, 99:
		condition = Thunderstorm
	default:
		condition = Unknown
	}

	mappedData := CurrentWeather{
		TempC:       convert.CelsiusToFahrenheit(finalData.Current.Temp),
		FeelsLikeC:  convert.CelsiusToFahrenheit(finalData.Current.FeelsLike),
		HumidityPct: finalData.Current.HumidityPct,
		WindKph:     convert.KphToMph(finalData.Current.WindKph),
		Condition:   condition,
	}

	finalForecast := Forecast{
		Current: mappedData,
		Hourly:  []HourPoint{},
		Daily:   []DayPoint{},
	}

	return finalForecast
}
