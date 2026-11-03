package main

import (
	"fmt"
	"os"

	"github.com/KillAllChickens/gust/internal/weather"
	"golang.org/x/term"
)

func TermWidth() int {
	for _, f := range []*os.File{os.Stdout, os.Stderr, os.Stdin} {
		fd := int(f.Fd())
		if term.IsTerminal(fd) {
			if w, _, err := term.GetSize(fd); err == nil && w > 0 {
				return w
			}
		}
	}
	return 80
}

func main() {
	forecast := weather.FetchForecast()
	width := TermWidth()
	msg := fmt.Sprintf("Temp: %f°F", forecast.Current.TempC)

	for _ = range (width/2) - (len(msg)/2) {
		fmt.Print(" ")
	}
	fmt.Print(msg)
	for _ = range (width/2) - (len(msg)/2) {
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
