# Sensor Metadata Application

## Task
Please build a JSON REST API for storing and querying sensor metadata.
At a minimum, this API should expose endpoints for the following:
- Storing name, location (gps position), and a list of tags for each sensor.
- Retrieving metadata for an individual sensor by name.
- Updating a sensor’s metadata.
- Querying to find the sensor nearest to a given location.

It is up to you how you structure your application, but please write it in Go and include anything you would
in a professional project (i.e.: README, tests, input validation, etc).


## API

### Endpoints

| METHOD | Endpoint |
| :----- | :------- |
| GET    | /sensors |
| POST   | /sensors |
| GET    | /sensors/:id |
| PATCH  | /update  |

### Sensor Object
```go
type sensor struct {
	Name     string   `json:"id"`
	Location location `json:"location"`
	Tags     []string `json:"tags,omitempty"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
```


## Status Codes

Sensor metadata returns the following status codes in its API:

| Status Code | Description |
| :-- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |