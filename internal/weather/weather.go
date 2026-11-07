package weather

import (
	"log"
	"time"

	"github.com/KillAllChickens/gust/internal/convert"
	"github.com/KillAllChickens/gust/internal/httpclient"
)

type respDataCurrent struct {
	Time        string  `json:"time"`
	Temp        float64 `json:"temperature_2m"`
	FeelsLike   float64 `json:"apparent_temperature"`
	HumidityPct int     `json:"relative_humidity_2m"`
	WindKph     float64 `json:"wind_speed_10m"`
	WeatherCode int     `json:"weather_code"`
}

type respDataHourly struct {
	Time       []string  `json:"time"`
	Temp       []float64 `json:"temperature_2m"`
	Code       []int     `json:"weather_code"`
	PrecipProb []*int    `json:"precipitation_probability"`
	Wind       []float64 `json:"wind_speed_10m"`
}

type respDataDaily struct {
	Time       []string  `json:"time"`
	TempMax    []float64 `json:"temperature_2m_max"`
	TempMin    []float64 `json:"temperature_2m_min"`
	Code       []int     `json:"weather_code"`
	PrecipProb []*int    `json:"precipitation_probability_max"`
	Sunrise    []string  `json:"sunrise"`
	Sunset     []string  `json:"sunset"`
}

type respData struct {
	Timezone string          `json:"timezone"`
	Current  respDataCurrent `json:"current"`
	Hourly   respDataHourly  `json:"hourly"`
	Daily    respDataDaily   `json:"daily"`
}

func FetchForecast() Forecast {
	lat, long := "37.9967", "-88.9201"
	// client := httpclient.Get()
	// client.SetBaseURL("https://api.open-meteo.com")
	// client.

	var finalData respData
	forecastResp, err := httpclient.Get().R().
		SetResult(&finalData).
		SetQueryParams(map[string]string{
			"current":   "temperature_2m,apparent_temperature,relative_humidity_2m,wind_speed_10m,weather_code",
			"hourly":    "temperature_2m,wind_speed_10m,weather_code,precipitation_probability",
			"daily":     "temperature_2m_max,temperature_2m_min,precipitation_probability_max,weather_code,sunrise,sunset",
			"timezone":  "auto",
			"latitude":  lat,
			"longitude": long,
		}).
		Get("https://api.open-meteo.com/v1/forecast")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	if forecastResp.StatusCode() >= 400 {
		log.Fatalf("bad status %d: %s", forecastResp.StatusCode(), forecastResp.String())
	}
	loc, _ := time.LoadLocation(finalData.Timezone)
	currentCondition := CodeToCondition(finalData.Current.WeatherCode)

	currentData := CurrentWeather{
		TempC:       convert.CelsiusToFahrenheit(finalData.Current.Temp),
		FeelsLikeC:  convert.CelsiusToFahrenheit(finalData.Current.FeelsLike),
		HumidityPct: finalData.Current.HumidityPct,
		WindKph:     convert.KphToMph(finalData.Current.WindKph),
		Condition:   currentCondition,
	}

	h := finalData.Hourly
	hourly := make([]HourPoint, 0, len(h.Time))
	for i, ts := range h.Time {
		t, _ := time.ParseInLocation(layoutDateTime, ts, loc)

		hourly = append(hourly, HourPoint{
			Time:         t,
			TempC:        convert.CelsiusToFahrenheit(h.Temp[i]),
			Condition:    CodeToCondition(h.Code[i]),
			PrecipChance: derefInt(h.PrecipProb[i]),
			WindKph:      convert.KphToMph(h.Wind[i]),
		})
	}

	d := finalData.Daily
	daily := make([]DayPoint, 0, len(d.Time))
	for i, ds := range d.Time {
		date, err := time.ParseInLocation(layoutDate, ds, loc)
		if err != nil {
			log.Fatalf("bad daily date %q: %v", ds, err)
		}
		sunrise, err := time.ParseInLocation(layoutDateTime, d.Sunrise[i], loc)
		if err != nil {
			log.Fatalf("bad sunrise %q: %v", d.Sunrise[i], err)
		}
		sunset, err := time.ParseInLocation(layoutDateTime, d.Sunset[i], loc)
		if err != nil {
			log.Fatalf("bad sunset %q: %v", d.Sunset[i], err)
		}
		daily = append(daily, DayPoint{
			Date:         date,
			HighC:        convert.CelsiusToFahrenheit(d.TempMax[i]),
			LowC:         convert.CelsiusToFahrenheit(d.TempMin[i]),
			Condition:    CodeToCondition(d.Code[i]),
			PrecipChance: derefInt(d.PrecipProb[i]),
			Sunrise:      sunrise,
			Sunset:       sunset,
		})
	}

	finalForecast := Forecast{
		Current: currentData,
		Hourly:  hourly,
		Daily:   daily,
	}

	return finalForecast
}

func CodeToCondition(code int) Condition {
	switch code {
	case 0, 1:
		return Clear
	case 2, 3:
		return Cloudy
	case 45, 48:
		return Fog
	case 51, 53, 55, 56, 57, 61, 63, 65, 66, 67, 80, 81, 82:
		return Rain
	case 71, 73, 75, 77, 85, 86:
		return Snow
	case 95, 96, 99:
		return Thunderstorm
	default:
		return Unknown
	}
}
