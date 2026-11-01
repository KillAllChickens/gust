package weather

import "time"

type Condition string

const (
	Clear        Condition = "clear"
	Cloudy       Condition = "cloudy"
	Rain         Condition = "rain"
	Snow         Condition = "snow"
	Thunderstorm Condition = "thunderstorm"
	Fog          Condition = "fog"
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
