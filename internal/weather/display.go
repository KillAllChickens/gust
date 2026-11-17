package weather

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KillAllChickens/gust/internal/config"
	"golang.org/x/term"
)

func ShowWeather() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config failed to load")
	}

	u := NewUnits(cfg.Units)
	forecast := FetchForecast(cfg.Lat, cfg.Lon)

	fmt.Printf("Its currently %s in %s\n", forecast.Current.Condition, cfg.Location)
	fmt.Printf("It is currently %.1f%s (feels like %.1f%s)\n",
		u.Temp(forecast.Current.TempC), u.TempUnit,
		u.Temp(forecast.Current.FeelsLikeC), u.TempUnit)
	fmt.Printf("Wind: %.0f %s", u.Wind(forecast.Current.WindKph), u.windUnit)

	now := time.Now()
	loc := now.Location()

	startOfHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, loc)
	startOfTomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, loc)

	var hours []HourPoint
	for _, h := range forecast.Hourly {
		if h.Time.Before(startOfHour) || !h.Time.Before(startOfTomorrow) {
			continue
		}
		hours = append(hours, h)
	}

	width := 80
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		width = w
	}

	fmt.Println()
	fmt.Println(RenderHourly(hours, width, 5, u))
}
