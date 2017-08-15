package main

import "fmt"

func main() {
	var a float64
	var c, d int32
	var e, f = 3.5, "Nazmul"
	var g string = "Basher"
	b:=4
	a = 5.67
	c =3
	d = 5
	e = 6
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)	
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("d is of type %T\n", d)
	fmt.Printf("e is of type %T\n", e)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
