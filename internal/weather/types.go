package weather

import (
	"time"

	"github.com/KillAllChickens/gust/internal/convert"
)

type Condition string

const (
	Clear        Condition = "clear"
	Cloudy       Condition = "cloudy"
	Rain         Condition = "rainy"
	Snow         Condition = "snowing"
	Thunderstorm Condition = "thunderstorming"
	Fog          Condition = "foggy"
	Unknown      Condition = "unknown"
)

type HourPoint struct {
	Time         time.Time
	TempC        float64
	Condition    Condition
	PrecipChance int
	WindKph      float64
}

type DayPoint struct {
	Date         time.Time
	HighC        float64
	LowC         float64
	Condition    Condition
	PrecipChance int
	Sunrise      time.Time
	Sunset       time.Time
}

type CurrentWeather struct {
	TempC       float64
	FeelsLikeC  float64
	Condition   Condition
	HumidityPct int
	WindKph     float64
}

type Forecast struct {
	Current CurrentWeather
	Hourly  []HourPoint
	Daily   []DayPoint
}

type Units struct {
	Temp     func(c float64) float64
	Wind     func(kph float64) float64
	TempUnit string
	windUnit string
}

func identity(v float64) float64 { return v }

func NewUnits(system string) Units {
	if system == "imperial" {
		return Units{
			Temp: convert.CelsiusToFahrenheit,
			Wind: convert.KphToMph,
			TempUnit: "°F",
			windUnit: "mph",
		}
	}
	return Units{identity, identity, "°C", "km/h"}
}

const (
	layoutDateTime = "2006-01-02T15:04"
	layoutDate     = "2006-01-02"
)

func derefInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
