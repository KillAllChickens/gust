package main

import (
	"fmt"

	"github.com/KillAllChickens/gust/internal/weather"
)

func main() {
	fmt.Printf("%+v\n", weather.FetchWeather())
}
