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

type Person struct {
	Name string
}

// Embedded Types example
type Android struct {
	Person
	Model string
}

var (
	s1 = Student{1, "Nazmul"}
	s2 = Student{Name: "Saiham"}
	s3 = Student{}
)

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

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
	
	// new function takes a type as an argument, allocates enough memory 
	// to fit a value of that type and returns a pointer to it.
	a := new(Android)
	a.Name = "Nazmul"
	a.Talk() // you can also use like a.Person.Talk()
}
