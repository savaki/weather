package weather

import (
	"fmt"
	"github.com/savaki/weather/openweathermap"
	"log"
	"testing"
	"time"
)

func ok(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func TestLive(t *testing.T) {
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

	results, err := service.FindAll(cities)
	ok(err)

	fmt.Println(results)
}
