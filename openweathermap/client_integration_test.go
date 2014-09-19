package openweathermap

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"code.google.com/p/go.net/context"
)

// +build integration

func TestIntegration(t *testing.T) {
	var c *client
	var result *WeatherResult
	var err error

	Convey("Given a weather client", t, func() {
		c = &client{}

		Convey("When I search for London", func() {
			result, err = c.GetWeather(context.Background(), "London")
		})

		Convey("Then I expect no errors", func() {
			So(err, ShouldBeNil)
		})

		Convey("And the temperate should be set", func() {
			So(result.Main.Temp, ShouldNotEqual, 0.0)
			So(result.Main.TempMax, ShouldNotEqual, 0.0)
			So(result.Main.TempMin, ShouldNotEqual, 0.0)
		})
	})
}
