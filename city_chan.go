package weather

import (
	"code.google.com/p/go.net/context"
	"errors"
	"github.com/savaki/weather/openweathermap"
	"fmt"
)

type RequestFunc func(context.Context, openweathermap.WeatherService) (*openweathermap.WeatherResult, error)

type response struct {
	id     int
	result *openweathermap.WeatherResult
	err    error
}

func (s *WeatherService) Find(requests <-chan RequestFunc) error {
	// cancel all things that we're done with
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	// internal communication channel
	responses := make(chan response)

	handle := func(id int, request RequestFunc) {
		i := id
		debug(fmt.Sprintf("request: %d", i))
		result, err := request(ctx, s.Client)
			responses <- response{
				id:     i,
				result: result,
				err:    err,
			}
	}

	// invoke a bunch of go-routines to make the call
	id := 0
	for request := range requests {
		id = id + 1
		go handle(id, request) // request #1
		go handle(id, request) // request #2
	}

	results := map[int]int{}
	for len(results) < id {
		select {
		case result := <-responses:
			if result.err == nil {
				results[result.id] = result.id
				debug(fmt.Sprintf("received - %d", result.id))
			}
		case <-ctx.Done():
			debug("timeout")
			return errors.New("must have timed out")
		}
	}

	debug("finished")

	return nil
}
