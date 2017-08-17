package main

import "fmt"

func main(){
	var b int = 8
	var a int = 2
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution */
	for a := 1; a < 5; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	
	fmt.Println(a)
	
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}

	fmt.Println(a)
	
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}
