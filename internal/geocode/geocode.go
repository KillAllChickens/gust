package geocode

import (
	"fmt"
	"log"

	"github.com/KillAllChickens/gust/internal/httpclient"
)

type CityResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	State       string `json:"admin1"`
	CountryCode string `json:"country_code"`
	Timezone    string `json:"timezone"`
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
		Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=3&language=en&format=json", query))
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range response.Results {
		fmt.Printf("%s, %s, %s (%s)\n", res.Name, res.State, res.CountryCode, res.Timezone)
	}
}
