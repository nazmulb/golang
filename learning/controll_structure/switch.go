package main

import "fmt"

func main() {
	var marks float32 = 40

	switch marks {
		case 90: 
			fmt.Println("Excellent")
		case 80:
			fmt.Println("Very Good")
		case 60:
			fmt.Println("Good")
		case 40, 50:
			fmt.Println("Not Bad")
		default:
			fmt.Println("Failed or Invalid input")
	}
}
