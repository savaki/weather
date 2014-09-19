package openweathermap

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// +build integration

func TestIntegration(t *testing.T) {
	var client *Client
	var result *WeatherResult
	var err error

	Convey("Given a weather client", t, func() {
		client = New()

		Convey("When I search for London", func() {
			result, err = client.GetWeather("London")
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
