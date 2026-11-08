package geocode

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/KillAllChickens/gust/internal/httpclient"
)

type CityResp struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"admin1"`
	CountryCode string  `json:"country_code"`
	Long        float64 `json:"latitude"`
	Lat         float64 `json:"longitude"`
	Timezone    string  `json:"timezone"`
}

type respData struct {
	Results []CityResp `json:"results"`
}

// curl "https://geocoding-api.open-meteo.com/v1/search?name=Austin&count=3&language=en&format=json"

func SearchCity(query string) {
	// geoClient := resty.New()
	// geoClient.SetHeader("User-Agent", "gust/1.0 (admin@killallchickens.org)")

	var response respData
	_, err := httpclient.Get().R().
		SetResult(&response).
		Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=10&language=en&format=json", query))
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tSTATE/REGION\tCOUNTRY\tTIMEZONE")
	for i, res := range response.Results {
		fmt.Fprintf(w, "[%d]\t%s\t%s\t%s\t%s\n", i+1, res.Name, res.State, res.CountryCode, res.Timezone)
	}
	w.Flush()
}
