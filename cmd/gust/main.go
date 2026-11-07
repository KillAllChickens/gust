package main

import (
	"github.com/KillAllChickens/gust/internal/geocode"
)

// func TermWidth() int {
// 	for _, f := range []*os.File{os.Stdout, os.Stderr, os.Stdin} {
// 		fd := int(f.Fd())
// 		if term.IsTerminal(fd) {
// 			if w, _, err := term.GetSize(fd); err == nil && w > 0 {
// 				return w
// 			}
// 		}
// 	}
// 	return 80
// }
//
// func centerPrint(width int, msg string) {
// 	for _ = range (width/2) - (len(msg)/2) {
// 		fmt.Print(" ")
// 	}
// 	fmt.Print(msg)
// 	for _ = range (width/2) - (len(msg)/2) {
// 		fmt.Print(" ")
// 	}
// 	fmt.Print("\n")
// }
//
// func main() {
// 	forecast := weather.FetchForecast()
// 	width := TermWidth()
// 	centerPrint(width, fmt.Sprintf("Temp: %f°F", forecast.Current.TempC))
//
//
// 	for _, hour := range forecast.Hourly {
// 		fmt.Printf("%s, %f\n", hour.Time, hour.TempC)
// 	}
// }

func main() {
	geocode.SearchCity("Benton")
}
