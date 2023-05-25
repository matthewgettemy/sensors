package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func printSensors() {
	for _, s := range sensors {
		fmt.Printf("%#v\n", s)
	}
	fmt.Println()
}

func TestDefaultResponse(t *testing.T) {
	router := gin.Default()
	router.GET("/", defaultHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	response, _ := ioutil.ReadAll(w.Body)
	mockResponse := `{"message":"Sensor Metadata Application"}`
	assert.Equal(t, string(response), mockResponse)
	assert.Equal(t, http.StatusOK, w.Code)
	printSensors()
}

func TestCreateNewSensor(t *testing.T) {
	router := gin.Default()
	router.POST("/sensors", createSensor)
	newSensor := sensor{
		Name:     "test_name",
		Location: location{Latitude: 90.0, Longitude: 91.0},
		Tags:     []string{"tag1", "tag2"},
	}
	j, _ := json.Marshal(newSensor)
	req, _ := http.NewRequest("POST", "/sensors", bytes.NewBuffer(j))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	printSensors()
}

func TestGetAllSensors(t *testing.T) {
	router := gin.Default()
	router.GET("/sensors", getSensors)
	req, _ := http.NewRequest("GET", "/sensors", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var retrievedSensors map[string]sensor
	err := json.Unmarshal(w.Body.Bytes(), &retrievedSensors)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, retrievedSensors)
}

func TestUpdateSensor(t *testing.T) {
	router := gin.Default()
	router.PUT("/sensor/:id", updateSensorById)
	newSensorId := uuid.New().String()
	newSensor := sensor{
		Name:     "test_name",
		Location: location{Latitude: 90.0, Longitude: 91.0},
		Tags:     []string{"tag1", "tag2"},
	}
	j, _ := json.Marshal(newSensor)
	reqNotFound, _ := http.NewRequest("PUT", "/sensor/"+newSensorId, bytes.NewBuffer(j))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
	sensors[newSensorId] = newSensor
	reqFound, _ := http.NewRequest("PUT", "/sensor/"+newSensorId, bytes.NewBuffer(j))
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, reqFound)
	assert.Equal(t, http.StatusOK, w2.Code)
}
