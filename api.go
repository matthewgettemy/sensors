package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getSensors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sensors)
}

func sensorById(c *gin.Context) {
	id := c.Param("id")
	sensor, err := getSensorById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sensor not found"})
	}

	c.IndentedJSON(http.StatusOK, sensor)
}

func createSensor(c *gin.Context) {
	var newSensor sensor

	if err := c.BindJSON(&newSensor); err != nil {
		return
	}

	sensors[uuid.New().String()] = newSensor
	c.IndentedJSON(http.StatusCreated, newSensor)
}

func closestSensorByLocation(c *gin.Context) {
	var loc location
	if err := c.BindJSON(loc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	id := closestSensorId(loc)
	closestSensor := sensors[id]
	c.IndentedJSON(http.StatusFound, closestSensor)
}

func updateSensorById(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
	}

	sensor, err := getSensorById(id)
	fmt.Printf("%+v", sensor)
	fmt.Printf("%s", err)
}
