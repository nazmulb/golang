package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Student struct {
	Id int
	Name string
}

var (
	s1 = Student{1, "Nazmul"}
	s2 = Student{Name: "Saiham"}
	s3 = Student{}
)

func (s Student) getName(pre string) string {
	return pre + " " + s.Name
}

func (s *Student) setName(name string) {
	s.Name = name
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	
	fmt.Println(s1.getName("Hello,"))
	s1.setName("Basher")
	fmt.Println(s1.getName("Basher"))
	fmt.Println(s1, s2, s3)	
}
