package main

import (
	"fmt"
	"testing"
)

func TestGetSensorById(t *testing.T) {
	addDummySensors(300)
	for id, sensor := range sensors {
		fmt.Printf("%s: %#v\n", id, sensor)
	}
}
