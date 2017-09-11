package main

import (
	"fmt"
	"github.com/nazmulb/golang/learning/nazmulpkg"
	"github.com/nazmulb/golang/learning/mathpkg"
)

func main() {
	fmt.Println(nazmulpkg.Show("Hello,"))

	xs := []float64{1,2,3,4}
	avg := mathpkg.Average(xs)
	fmt.Println(avg)
}
