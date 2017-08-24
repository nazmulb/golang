package main

import "fmt"

func main() {
	var marks float32 = 85
	var grade string

	switch {
		case marks >= 90 && marks <= 100: 
			grade = "A+"
		case marks >= 80 && marks < 90:
			grade = "A"
		case marks >= 60 && marks < 80:
			grade = "B+"
		case marks >= 33 && marks < 60:
			grade = "B"
		default:
			grade = "F"
	}


	switch grade {
		case "A+":
			fmt.Println("Excellent")
		case "A":
			fmt.Println("Very Good")
		case "B+", "B":
			fmt.Println("Good")
		default:
			fmt.Println("Failed or invalid")	
	}
}
