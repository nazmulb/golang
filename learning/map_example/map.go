package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var n = map[string]string{
	"BD": "Bangladesh",
	"USA": "United State of Amiraca",
}

var p = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	n["IN"] = "India"
	n["UK"] = "United Kingdom"	
	fmt.Println(n)

	delete(n, "IN")
	fmt.Println("The value:", n["IN"])

	v, ok := n["USA"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println(p)
}
