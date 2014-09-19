package openweathermap

import (
	"code.google.com/p/go.net/context"
	"encoding/json"
	. "github.com/visionmedia/go-debug"
	"net/http"
)

var debug = Debug("openweathermap")

type WeatherService interface {
	GetWeather(ctx context.Context, url string) (*WeatherResult, error)
}

func New() WeatherService {
	return &client{}
}

type client struct {
}

func (c *client) GetWeather(ctx context.Context, city string) (*WeatherResult, error) {
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?q="+city, nil)
	if err != nil {
		return nil, err
	}

	var result WeatherResult
	err = httpDo(ctx, req, func(response *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer response.Body.Close()

		return json.NewDecoder(response.Body).Decode(&result)
	})

	return &result, err
}

type Weather struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Id          int    `json:"id"`
	Main        string `json:"main"`
}

type Main struct {
	Temp     float32 `json:"temp"`
	TempMax  float32 `json:"temp_max"`
	TempMin  float32 `json:"temp_min"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type WeatherResult struct {
	Clouds  map[string]int `json:"clouds"`
	Id      int            `json:"id"`
	Main    Main           `json:"main"`
	Weather []Weather      `json:"weather"`
}
