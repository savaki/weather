package openweathermap

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestUnmarshalJson(t *testing.T) {
	var sample string
	var result WeatherResult
	var err error

	Convey("Given sample json content", t, func() {
		sample = `
{"coord":{"lon":-0.13,"lat":51.51},"sys":{"type":1,"id":5091,"message":0.2037,"country":"GB","sunrise":1411105356,"sunset":1411149938},"weather":[{"id":520,"main":"Rain","description":"light intensity shower rain","icon":"09n"},{"id":701,"main":"Mist","description":"mist","icon":"50n"},{"id":741,"main":"Fog","description":"fog","icon":"50n"}],"base":"cmc stations","main":{"temp":289.49,"pressure":1010,"humidity":100,"temp_min":288.15,"temp_max":290.55},"wind":{"speed":2.1,"deg":340,"var_beg":260,"var_end":50},"rain":{"1h":3.18},"clouds":{"all":24},"dt":1411093488,"id":2643743,"name":"London","cod":200}
`

		Convey("When I decode this to a WeatherResult", func() {
			err = json.NewDecoder(strings.NewReader(sample)).Decode(&result)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
