package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
	var sensors = []sensor{
		{ID: "1", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
		{ID: "2", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
		{ID: "3", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
		{ID: "4", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
	}
*/
var sensors = map[string]sensor{
	uuid.New().String(): sensor{Name: "sensor1", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
	uuid.New().String(): sensor{Name: "sensor2", Location: location{Latitude: 1.0, Longitude: 1.0}, Tags: []string{"hi", "wut"}},
	uuid.New().String(): sensor{Name: "sensor3", Location: location{Latitude: 2.0, Longitude: 2.0}, Tags: []string{"hi", "bye"}},
}

func main() {
	router := gin.Default()
	router.POST("/sensors", createSensor)
	router.GET("/sensors", getSensors)
	router.GET("/sensors/:id", sensorById)
	router.GET("/sensors/closest", closestSensorByLocation)
	router.PATCH("/update", updateSensorById)
	router.Run("localhost:8080")
}
