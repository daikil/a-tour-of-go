package main

import "fmt"

type Vertex3 struct {
	Lat, Long float64
}

var m3 = map[string]Vertex3{
	"Bell Labs": {40.684, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m3)
}
