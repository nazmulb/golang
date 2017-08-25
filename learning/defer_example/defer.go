package main

import "fmt"

func main() {
	defer fmt.Println("Nazmul")

	fmt.Println("Hello")


	fmt.Println("Start")
	
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("End")
}
