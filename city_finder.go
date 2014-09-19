package weather

import (
	"code.google.com/p/go.net/context"
	"errors"
	"fmt"
	"github.com/savaki/weather/openweathermap"
	. "github.com/visionmedia/go-debug"
	"time"
)

var debug = Debug("weather")

type WeatherService struct {
	Client  openweathermap.WeatherService
	Timeout time.Duration
}

type getWeatherResponse struct {
	city    string
	weather *openweathermap.WeatherResult
	err     error
}

func (s *WeatherService) FindAll(cities []string) (map[string]*openweathermap.WeatherResult, error) {
	// cancel all things that we're done with
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	// internal communication channel
	responses := make(chan *getWeatherResponse)

	// invoke a bunch of go-routines to make the call
	for _, city := range cities {
		_city := city
		go s.find(ctx, _city, responses) // call #1
		go s.find(ctx, _city, responses) // call #2
	}

	// collect the responses; return once we've gathered a result for each city
	allWeather := map[string]*openweathermap.WeatherResult{}
	for cityCount := len(cities); len(allWeather) < cityCount; {
		select {
		case result := <-responses:
			if result.err == nil {
				debug(fmt.Sprintf("allWeather => %d", len(allWeather)))
				allWeather[result.city] = result.weather
			}
		case <-ctx.Done():
			debug("timeout")
			return nil, errors.New("must have timed out")
		}
	}

	return allWeather, nil
}

func (s *WeatherService) find(ctx context.Context, city string, responses chan *getWeatherResponse) {
	finder := func() *getWeatherResponse {
		response, err := s.Client.GetWeather(ctx, city)
		debug(city)
		return &getWeatherResponse{
			city:    city,
			weather: response,
			err:     err,
		}
	}

	select {
	case responses <- finder():
	case <-ctx.Done():
	}
}
