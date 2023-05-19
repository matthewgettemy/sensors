package main

import (
	"errors"
)

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

func getSensorById(id string) (*sensor, error) {
	targetSensor, ok := sensors[id]
	if !ok {
		return nil, errors.New("sensor not found")
	}
	return &targetSensor, nil
}

func updateSensor() {

}
