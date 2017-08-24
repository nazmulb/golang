package main

import "fmt"

func main() {
	var a, b int = 3, 5
	var i int = 10
	var j int = 20
	
	x, y := swap(a, b)
	fmt.Printf("x: %d, y: %d\n", x, y)
	_, z := swap(x, y)
	fmt.Printf("z: %d\n", z)

	swapByRef(&i, &j)
	fmt.Printf("i: %d, j: %d\n", i, j)	
	
	fmt.Println(add(x, y))

	getName := func(str string) string {
		return str
	}

	fmt.Println(getName("Nazmul"))

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	
	add := func(a, b int) int {
		return a + b	
	}
	
	fmt.Println(calculator(add))
	
	fmt.Println(calculator(func(a, b int) int {
		return b-a
	}))
}

func swap(a, b int) (int, int) {
	return b, a
}

func swapByRef(a *int, b *int) {
	var temp int
   	temp = *a    /* save the value at address a */
   	*a = *b    /* put b into a */
   	*b = temp    /* put temp into b */
}

func add(a int, b int) int {
	return a+b
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func calculator(cb func(int, int) int) int {
	a := 10
	b := 20	
	
	return cb(a, b)
}
