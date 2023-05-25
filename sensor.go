package main

import (
	"errors"
	"math"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
	"github.com/matthewgettemy/sensors/cities"
	"github.com/umahmood/haversine"
)

const (
	temperature string = "temperature"
	pressure    string = "pressure"
	vibration   string = "vibration"
)

var sensorTypes = []string{temperature, pressure, vibration}

// JSON encoder can only see, and therefore only encode, exported fields in a struct
type sensor struct {
	Name     string   `json:"id"`
	Location location `json:"location"`
	Tags     []string `json:"tags,omitempty"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (loc location) coords() haversine.Coord {
	return haversine.Coord{Lat: loc.Latitude, Lon: loc.Longitude}
}

func getSensorById(id string) (*sensor, error) {
	targetSensor, ok := sensors[id]
	if !ok {
		return nil, errors.New("sensor not found")
	}
	return &targetSensor, nil
}

func closestSensorId(loc location, metric bool) (string, float64) {
	minDist := math.Inf(0)
	var closest string
	for id, sensor := range sensors {
		dist := haversineDistance(loc, sensor.Location, metric)
		if dist < minDist {
			minDist = dist
			closest = id
		}
	}
	return closest, minDist
}

func haversineDistance(location1 location, location2 location, metric bool) float64 {
	/*
		Get the great-circle distance between two points on a sphere given lat and lon.
		Haversine formula: a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
						   c = 2 ⋅ atan2(√a, √(1−a))
						   d = R ⋅ c
			where φ is latitude, λ is longitude, R is earth’s radius (mean radius = 6,371km)
		Latitude: -90 -> 90
		Longitude: -180 -> 180
	*/
	loc1Coords := location1.coords()
	loc2Coords := location2.coords()
	mi, km := haversine.Distance(loc1Coords, loc2Coords)
	var distance float64
	if metric {
		distance = km
	} else {
		distance = mi
	}
	return distance
}

func addDummySensors(numberToAdd int) {
	cities := cities.GetCityData("cities/cities.json")
	for i := 1; i <= numberToAdd; i++ {
		randomCity := cities[rand.Intn(len(cities))]
		randomType := sensorTypes[rand.Intn(len(sensorTypes))]
		// fmt.Printf("Adding %s sensor at city %s.\n", randomType, randomCity.City)
		newSensor := sensor{
			Name:     randomType + strconv.Itoa(i),
			Location: location{Latitude: randomCity.Latitude, Longitude: randomCity.Longitude},
			Tags:     []string{"tag1", "tag2"},
		}
		sensors[uuid.New().String()] = newSensor
	}
}
