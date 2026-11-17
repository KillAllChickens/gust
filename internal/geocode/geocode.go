package geocode

import (
	"log"
	"strings"

	"github.com/KillAllChickens/gust/internal/httpclient"
)

type CityResp struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"admin1"`
	County      string  `json:"admin2"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Long        float64 `json:"longitude"`
	Lat         float64 `json:"latitude"`
	Timezone    string  `json:"timezone"`
}

type RespData struct {
	Results []CityResp `json:"results"`
}

func fetch(name string) []CityResp {
	var response RespData
	_, err := httpclient.Get().R().
		SetResult(&response).
		SetQueryParams(map[string]string{
			"name":     name,
			"count":    "100",
			"language": "en",
			"format":   "json",
		}).
		Get("https://geocoding-api.open-meteo.com/v1/search")
	if err != nil {
		log.Fatal(err)
	}
	return response.Results
}

func matches(c CityResp, terms []string) bool {
	hay := strings.ToLower(strings.Join(
		[]string{c.State, c.County, c.Country, c.CountryCode}, " "))
	for _, t := range terms {
		if !strings.Contains(hay, strings.ToLower(t)) {
			return false
		}
	}
	return true
}

func SearchCity(query string) RespData {
	tokens := strings.Fields(strings.ReplaceAll(query, ",", " "))

	for i := len(tokens); i >= 1; i-- {
		name := strings.Join(tokens[:i], " ")
		terms := tokens[i:]

		var out []CityResp
		for _, c := range fetch(name) {
			if matches(c, terms) {
				out = append(out, c)
			}
		}
		if len(out) > 0 {
			return RespData{Results: out[:min(len(out), 10)]}
		}
	}
	return RespData{}
}
