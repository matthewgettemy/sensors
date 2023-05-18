package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sensors = []sensor{
	{ID: "1", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
	{ID: "2", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
	{ID: "3", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
	{ID: "4", Location: location{Latitude: 0.0, Longitude: 0.0}, Tags: []string{"hi"}},
}

func getSensors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sensors)
}

func getSensorById(id string) (*sensor, error) {
	for i, s := range sensors {
		if s.ID == id {
			return &sensors[i], nil
		}
	}

	return nil, errors.New("sensor not found")
}

func sensorById(c *gin.Context) {
	id := c.Param("id")
	sensor, err := getSensorById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sensor not found."})
	}

	c.IndentedJSON(http.StatusOK, sensor)
}

func createSensor(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Creating new sensor."})
}

func updateSensor(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
	}

	sensor, err := getSensorById(id)
	fmt.Printf("%+v", sensor)
	fmt.Printf("%s", err)
}

func main() {
	router := gin.Default()
	router.POST("/sensors", createSensor)
	router.GET("/sensors", getSensors)
	router.GET("/sensors/:id", sensorById)
	router.PATCH("/update", updateSensor)
	router.Run("localhost:8080")
}
