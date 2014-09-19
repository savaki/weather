package openweathermap

import (
	"code.google.com/p/go.net/context"
	"errors"
	"math/rand"
	"time"
)

var (
	generator = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// create a client that delays for a random amount up millis up to the specified delay
func Mock(delay int64, result *WeatherResult, err error) WeatherService {
	return &mock{
		delay:  time.Duration(generator.Int63n(delay)),
		result: result,
		err:    err,
	}
}

type mock struct {
	delay  time.Duration
	result *WeatherResult
	err    error
}

func (m *mock) GetWeather(ctx context.Context, city string) (*WeatherResult, error) {
	timer := time.NewTimer(m.delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, errors.New("operation canceled")
	case <-timer.C:
		return m.result, m.err
	}
}
