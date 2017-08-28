package main

import "fmt"

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return float64(v.X + v.Y)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var a Abser
	f := MyFloat(2.5)

	a = f
	describe(a)
	fmt.Println(a.Abs())

	v := Vertex{3.4, 4.2}

	a = &v
	describe(a)
	fmt.Println(a.Abs())

	var i interface{}
	describe(i)

	i = 4
	describe(i)

	i = "Nazmul"
	describe(i)

	x := Person{"Nazmul Basher", 36}
	y := Person{"Nabil Al Noor", 7}

	fmt.Println(x, y)
}
