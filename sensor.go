package main

type sensor struct {
	ID       string   `json:"id"`
	Location location `json:"location"`
	Tags     []string `json:"tags"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
