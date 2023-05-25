package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func defaultHandler(c *gin.Context) {
	// GET /
	c.JSON(http.StatusOK, gin.H{"message": "Sensor Metadata Application"})
}

func getSensors(c *gin.Context) {
	// GET /sensors
	c.JSON(http.StatusOK, sensors)
}

func sensorById(c *gin.Context) {
	// GET /sensors/:id
	id := c.Param("id")
	sensor, err := getSensorById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "sensor not found"})
	}

	c.JSON(http.StatusOK, sensor)
}

func createSensor(c *gin.Context) {
	// POST /sensors
	var newSensor sensor

	if err := c.BindJSON(&newSensor); err != nil {
		return
	}
	newId := uuid.New().String()
	sensors[newId] = newSensor
	c.JSON(http.StatusCreated, newId)
}

func closestSensorByLocation(c *gin.Context) {
	// POST /sensors/closest
	var loc location
	if err := c.BindJSON(&loc); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	metric := true
	id, dist := closestSensorId(loc, metric)
	fmt.Printf("Closest sensor %s is %f km away.\n", id, dist)
	closestSensor := sensors[id]
	c.JSON(http.StatusFound, closestSensor)
}

func updateSensorById(c *gin.Context) {
	// PUT /update/:id
	id := c.Param("id")
	_, err := getSensorById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "sensor not found"})
		return
	}

	var newSensor sensor
	if err := c.ShouldBindJSON(&newSensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sensors[id] = newSensor
	c.JSON(http.StatusOK, newSensor)
}
