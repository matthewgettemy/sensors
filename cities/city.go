package cities

import (
	"encoding/json"
	"io/ioutil"
)

type city struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetCityData(file string) []city {
	J, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var cities []city
	err = json.Unmarshal(J, &cities)
	if err != nil {
		panic(err)
	}
	return cities
}
