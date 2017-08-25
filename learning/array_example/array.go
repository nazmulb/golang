package main

import "fmt"

func main() {
	var a [2]string

	 b := [6]int{2, 3, 4, 5, 6, 77}

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var c []int = b[1:4]
	fmt.Println(c)
}
