package main

import (
	"fmt"
	"testing"

	"github.com/matthewgettemy/sensors/cities"
)

func TestGetSensorById(t *testing.T) {
	sensors = make(map[string]sensor)
	addDummySensors(10)
	testId := "dummy_id"
	testSensor := sensor{Name: "test_name", Location: location{Latitude: 100.0, Longitude: 200.0}, Tags: []string{"hi", "bye"}}
	sensors[testId] = testSensor
	sensorFound, err := getSensorById(testId)
	if err != nil {
		t.Errorf("could not find sensor %v#", testSensor)
	}
	if sensorFound.Name != testSensor.Name {
		t.Errorf("Got %#v want %#v.", sensorFound, testSensor)
	}
	/*
		for id, sensor := range sensors {
			fmt.Printf("%s: %#v\n", id, sensor)
		}
	*/
}

func TestClosestSensorId(t *testing.T) {
	sensors = make(map[string]sensor)
	addDummySensors(500)
	cities := cities.GetCityData("cities/cities.json")
	targetCity := cities[0]
	targetLocation := location{Latitude: targetCity.Latitude, Longitude: targetCity.Longitude}
	fmt.Println(targetCity.City)

	testLocation := location{Latitude: targetLocation.Latitude + 0.01, Longitude: targetLocation.Longitude + 0.01}
	testId := "dummy_id"
	sensors[testId] = sensor{Name: "test_name", Location: location{Latitude: targetCity.Latitude, Longitude: targetCity.Longitude}, Tags: []string{"hi"}}
	foundId, _ := closestSensorId(testLocation, true)
	closestSensor := sensors[foundId]
	if (targetLocation.Latitude != closestSensor.Location.Latitude) && (targetLocation.Longitude != closestSensor.Location.Longitude) {
		t.Errorf("Looking for location %#v, found location %#v.", targetLocation, closestSensor.Location)
	}
}

func TestOneSensorClosest(t *testing.T) {
	sensors = make(map[string]sensor)
	addDummySensors(1)
	testLocation := location{Latitude: 50.1, Longitude: 100.1}
	foundId, dist := closestSensorId(testLocation, true)
	fmt.Printf("id=%s, dist=%f km\n", foundId, dist)
}
