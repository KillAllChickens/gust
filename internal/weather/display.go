package weather

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/KillAllChickens/gust/internal/config"
	"golang.org/x/term"
)

var termWidth int

func init() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	termWidth = width
}

var ansiRe = regexp.MustCompile(`\x1b\[[0-9;?]*[ -/]*[@-~]`)

func visibleLen(s string) int {
	return utf8.RuneCountInString(ansiRe.ReplaceAllString(s, ""))
}

func centerPrintf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	pad := (termWidth - visibleLen(s)) / 2
	if pad < 0 {
		pad = 0
	}
	fmt.Print(strings.Repeat(" ", pad) + s)
}

func ShowWeather() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config failed to load")
	}

	u := NewUnits(cfg.Units)
	fmt.Println("Fetching weather for update...")
	forecast := FetchForecast(cfg.Lat, cfg.Lon)

	fmt.Print("\033[H\033[2J") // clear

	centerPrintf("\033[1m%s\033[0m\n\n", time.Now().Format("3:04 PM"))
	centerPrintf("Its currently %s in %s\n\n", forecast.Current.Condition, cfg.Location)
	centerPrintf("It is currently %.1f%s (feels like %.1f%s)\n\n",
		u.Temp(forecast.Current.TempC), u.TempUnit,
		u.Temp(forecast.Current.FeelsLikeC), u.TempUnit)
	centerPrintf("Wind: %.0f %s\n", u.Wind(forecast.Current.WindKph), u.windUnit)

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
