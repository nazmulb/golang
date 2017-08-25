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
	fmt.Printf("length: %d, capacity %d\n", len(c), cap(c))

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
	
	x := make([]int, 5)
	fmt.Println(x)
	
	var y []int

	y = append(y, 2, 3, 7)
	fmt.Println(y)

	y = append(y, 10)
	fmt.Println(y)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("%d = %d\n", i, v)
	}
}
