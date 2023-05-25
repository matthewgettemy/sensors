package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// initialize some sensors
var sensors = map[string]sensor{
	uuid.New().String(): {Name: "sensor1", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
	uuid.New().String(): {Name: "sensor2", Location: location{Latitude: 1.0, Longitude: 1.0}, Tags: []string{"hi", "wut"}},
	uuid.New().String(): {Name: "sensor3", Location: location{Latitude: 2.0, Longitude: 2.0}, Tags: []string{"hi", "bye"}},
}

func main() {
	router := gin.Default()
	router.GET("/", defaultHandler)
	router.POST("/sensors", createSensor)
	router.GET("/sensors", getSensors)
	router.GET("/sensors/:id", sensorById)
	router.GET("/sensors/closest", closestSensorByLocation)
	router.PUT("/update/:id", updateSensorById)
	router.Run("localhost:8080")
}
