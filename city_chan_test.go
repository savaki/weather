package weather

import (
	"code.google.com/p/go.net/context"
	"github.com/savaki/weather/openweathermap"
	"testing"
	"time"
)

func CityChan(cities []string) <-chan RequestFunc {
	results := make(chan RequestFunc)

	go func() {
		for _, city := range cities {
			_city := city
			results <- func(ctx context.Context, client openweathermap.WeatherService) (*openweathermap.WeatherResult, error) {
				debug(_city)
				return client.GetWeather(ctx, "http://api.openweathermap.org/data/2.5/weather?q="+_city)
			}
		}

		close(results)
	}()


	return (<-chan RequestFunc)(results)
}

func TestFindFunc(t *testing.T) {
	service := WeatherService{
		Client:  openweathermap.New(),
		Timeout: time.Second * 15,
	}

	cities := []string{
		"San Francisco",
		"Oakland",
		"Detroit",
		"London",
	}
	requests := CityChan(cities)

	service.Find(requests)
	debug("done")
}
