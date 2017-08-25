package main

import "fmt"

func main() {
	var a [2]string

	 b := [6]int{2, 3, 4, 5, 6, 77}
	d := []int{6, 78, 44, 5, 7}	

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var c []int = b[1:4]
	fmt.Println(c)

	fmt.Println(d)

	s := []struct {
		id int
		name string
	}{
		{1, "Nazmul"},
		{2, "Saiham"},
		{3, "Abu Sayed"},
	}

	fmt.Println(s)

}
