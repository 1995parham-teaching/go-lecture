package main

import "fmt"

type Student struct {
	Name   string
	Family string
	age    int
}

// New is a popular pattern to create types
func New(name, family string, age int) Student {
	return Student{
		Name:   name,
		Family: family,
		age:    age,
	}
}

func (Student) Nothing() {}

func (Student) nothing() {}

// String returns the string for printing by %v or Println.
func (s Student) String() string {
	return fmt.Sprintf("Name: %s, Family: %s, age: %d", s.Name, s.Family, s.age)
}

func (s Student) Hello(to Student) string {
	return fmt.Sprintf("Hello %s, I am %s %s (%d)", to.Name, s.Name, s.Family, s.age)
}

func main() {
	s := Student{
		Name:   "Parham",
		Family: "Alvani",
		age:    27,
	}

	fmt.Println(s)

	s.nothing()

	fmt.Printf("student, %s %s\n", s.Name, s.Family)

	fmt.Println(s.Hello(Student{
		Name: "Torvalds",
	}))
}
