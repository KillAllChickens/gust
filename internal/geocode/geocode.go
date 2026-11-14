package geocode

import (
	"fmt"
	"log"
	"net/url"

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

type RespData struct {
	Results []CityResp `json:"results"`
}

// curl "https://geocoding-api.open-meteo.com/v1/search?name=Austin&count=3&language=en&format=json"

func SearchCity(query string) RespData {
	// geoClient := resty.New()
	// geoClient.SetHeader("User-Agent", "gust/1.0 (admin@killallchickens.org)")

	var response RespData
	_, err := httpclient.Get().R().
		SetResult(&response).
		Get(fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=10&language=en&format=json", url.PathEscape(query)))
	if err != nil {
		log.Fatal(err)
	}

	return response
}
